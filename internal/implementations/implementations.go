package impls

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/output"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/container/intsets"
	"golang.org/x/tools/go/packages"
)

func AddImplementationRelationships(
	pkgs []*packages.Package,
	// documents map[string]*document.Document,
	symbols *lookup.Global,
) {

	output.WithProgress("Implementations", func() {
		localInterfaces, localTypes, err := extractInterfacesAndConcreteTypes(pkgs, symbols)
		if err != nil {
			return
		}

		// Create a unique mapping of method -> int
		//	Then we'll use sparse sets to lookup whether things duck type or not
		allMethods := map[string]int{}

		for _, iface := range localInterfaces {
			for _, method := range iface.Methods {
				if _, ok := allMethods[method]; !ok {
					allMethods[method] = len(allMethods)
				}
			}
		}

		for _, ty := range localTypes {
			for _, method := range ty.Methods {
				if _, ok := allMethods[method]; !ok {
					allMethods[method] = len(allMethods)
				}
			}
		}

		// Create a map of method names to corresponding interfaces
		interfaceToMethodSet := map[string]*intsets.Sparse{}
		for _, iface := range localInterfaces {
			if iface.Ident == nil {
				continue
			}

			methodSet := &intsets.Sparse{}
			for _, method := range iface.Methods {
				methodSet.Insert(allMethods[method])
			}

			interfaceToMethodSet[iface.Symbol] = methodSet
		}

		fmt.Println("Types:")
		for _, ty := range localTypes {
			pos := ty.Ident.Pos()
			sym, ok := symbols.GetSymbolInformation(ty.Pkg, pos)
			if !ok {
				panic(fmt.Sprintf("Could not find symbol for %s", ty.Symbol))
			}

			fmt.Println(ty.Symbol, sym)

			methodSet := &intsets.Sparse{}
			for _, method := range ty.Methods {
				methodSet.Insert(allMethods[method])
			}

			fmt.Println("  Starting: ", ty.Ident.Name, methodSet)
			tyImpls := implementationsForType(ty, methodSet, interfaceToMethodSet)
			fmt.Println("  => ", tyImpls)
			for _, impl := range tyImpls {
				sym.Relationships = append(sym.Relationships, &scip.Relationship{
					Symbol:           impl,
					IsReference:      false,
					IsImplementation: true,
					IsTypeDefinition: false,
					IsDefinition:     false,
				})
			}

		}
	})
}

func implementationsForType(ty ImplDef, tyMethods *intsets.Sparse, interfaceToMethodSet map[string]*intsets.Sparse) (matching []string) {
	// Empty type - skip it.
	if len(ty.Methods) == 0 {
		return
	}

	for symbol, methods := range interfaceToMethodSet {
		if methods.SubsetOf(tyMethods) {
			matching = append(matching, symbol)
		}
	}

	return matching

	// // Find all the concrete types that implement this interface.
	// // Types that implement this interface are the intersection
	// // of all sets of receivers of all methods in this interface.
	// candidateTypes := &intsets.Sparse{}
	//
	// // The rest of this function is effectively "fold" (for those CS PhDs out there).
	// //
	// // > I think the underlying logic here is really beautiful but the syntax
	// // > makes it a bit messy and really obscures the intent and simplicity
	// // > behind it
	// //
	// //    - Dr. Fritz
	//
	// // If it doesn't match on the first method, then we can immediately quit.
	// // Concrete types must _always_ implement all the methods
	// if initialReceivers, ok := methodsToInterfaces[ty.Methods[0]]; !ok {
	// 	return
	// } else {
	// 	candidateTypes.Copy(initialReceivers)
	// }
	//
	// // Loop over the rest of the methods and find all the types that intersect
	// // every method of the interface.
	// for _, method := range ty.Methods[1:] {
	// 	receivers, ok := methodsToInterfaces[method]
	// 	if !ok {
	// 		return
	// 	}
	//
	// 	candidateTypes.IntersectionWith(receivers)
	// 	if candidateTypes.IsEmpty() {
	// 		return
	// 	}
	// }
	//
	// // Add the implementations to the relation.
	//
	// for _, ty := range candidateTypes.AppendTo(nil) {
	// 	matching = append(matching, interfaceToMethodSet[ty].Ident.Name)
	// }
	//
	// return
}

// monikerPackage:     monikerPackage,
// monikerIdentifier:  joinMonikerParts(monikerPackage, makeMonikerIdentifier(i.packageDataCache, pkg, obj)),
// typeNameIsExported: typeName.Exported(),
// typeNameIsAlias:    typeName.IsAlias(),
// identIsExported:    ident.IsExported(),
// defInfo:            i.getDefinitionInfo(typeName, ident),
// methods:            canonicalizedMethods,
// methodsByName:      methodsByName,
type ImplDef struct {
	// The corresponding scip symbol, generated via previous iteration over the AST
	Symbol string

	Pkg             *packages.Package
	Ident           *ast.Ident
	IsExportedIdent bool
	IsExportedType  bool
	IsAliasType     bool
	Methods         []string
}

func extractInterfacesAndConcreteTypes(pkgs []*packages.Package, symbols *lookup.Global) (interfaces []ImplDef, concreteTypes []ImplDef, err error) {
	for _, pkg := range pkgs {
		// Builtin isn't the same as standard library, that is for builtin types
		// We don't need to check those for implemenations.
		if pkg.Name == "builtin" {
			continue
		}

		if pkg == nil || pkg.TypesInfo == nil {
			panic(fmt.Sprintf("nill types info %s", pkg.Name))
		}

		pkgSymbols := symbols.GetPackage(pkg)
		if pkgSymbols == nil {
			fmt.Println("No symbols for package:", pkg.Name)
			continue
		}

		for ident, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}

			// fmt.Printf("extracting: %s %s %T\n", ident.Name, obj.Name(), obj)

			// We ignore aliases 'type M = N' to avoid duplicate reporting
			// of the Named type N.
			obj, ok := obj.(*types.TypeName)
			if !ok {
				continue
			}

			objType, ok := obj.Type().(*types.Named)
			if !ok {
				continue
			}

			symbol, ok := pkgSymbols.Get(obj.Pos())
			if !ok {
				fmt.Println("No symbol for:", ident.Name, obj.Name())
				continue
			}

			methods := listMethods(objType)

			// ignore interfaces that are empty. they are too
			// plentiful and don't provide useful intelligence.
			if len(methods) == 0 {
				continue
			}

			canonicalizedMethods := []string{}
			for _, m := range methods {
				canonicalizedMethods = append(canonicalizedMethods, canonicalizeMethod(m))
			}

			d := ImplDef{
				Symbol:          symbol.Symbol,
				Pkg:             pkg,
				Ident:           ident,
				IsExportedIdent: obj.Exported(),
				IsExportedType:  obj.Exported(),
				IsAliasType:     obj.IsAlias(),
				Methods:         canonicalizedMethods,
			}

			if types.IsInterface(objType) {
				interfaces = append(interfaces, d)
			} else {
				concreteTypes = append(concreteTypes, d)
			}

		}
	}

	return
}

// listMethods returns the method set for a named type T
// merged with all the methods of *T that have different names than
// the methods of T.
//
// Copied from https://github.com/golang/tools/blob/1a7ca93429f83e087f7d44d35c0e9ea088fc722e/cmd/godex/print.go#L355
func listMethods(T *types.Named) []*types.Selection {
	// method set for T
	mset := types.NewMethodSet(T)
	var res []*types.Selection
	for i, n := 0, mset.Len(); i < n; i++ {
		res = append(res, mset.At(i))
	}

	// add all *T methods with names different from T methods
	pmset := types.NewMethodSet(types.NewPointer(T))
	for i, n := 0, pmset.Len(); i < n; i++ {
		pm := pmset.At(i)
		if obj := pm.Obj(); mset.Lookup(obj.Pkg(), obj.Name()) == nil {
			res = append(res, pm)
		}
	}

	return res
}

// Returns a string representation of a method that can be used as a key for finding matches in interfaces.
func canonicalizeMethod(m *types.Selection) string {
	builder := strings.Builder{}

	writeTuple := func(t *types.Tuple) {
		for i := 0; i < t.Len(); i++ {
			builder.WriteString(t.At(i).Type().String())
		}
	}

	signature := m.Type().(*types.Signature)

	// if an object is not exported, then we need to make the canonical
	// representation of the object not able to match any other representations
	if !m.Obj().Exported() {
		builder.WriteString(m.Obj().Pkg().Path())
		builder.WriteString(":")
	}

	builder.WriteString(m.Obj().Name())
	builder.WriteString("(")
	writeTuple(signature.Params())
	builder.WriteString(")")

	returnTypes := signature.Results()
	returnLen := returnTypes.Len()
	if returnLen == 0 {
		// Don't add anything
	} else if returnLen == 1 {
		builder.WriteString(" ")
		writeTuple(returnTypes)
	} else {
		builder.WriteString(" (")
		writeTuple(returnTypes)
		builder.WriteString(")")
	}

	// fmt.Println(builder.String())
	return builder.String()
}
