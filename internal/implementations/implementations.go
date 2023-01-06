package impls

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/output"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/container/intsets"
	"golang.org/x/tools/go/packages"
)

type canonicalMethod string
type scipSymbol string

type ImplDef struct {
	// The corresponding scip symbol, generated via previous iteration over the AST
	Symbol *scip.SymbolInformation

	Pkg     *packages.Package
	Ident   *ast.Ident
	Methods map[canonicalMethod]*scip.SymbolInformation

	// IsExportedIdent bool
	// IsExportedType  bool
	// IsAliasType     bool
}

func findImplementations(concreteTypes map[string]ImplDef, interfaces map[string]ImplDef, symbols *lookup.Global) {
	// Create a unique mapping of method -> int
	//	Then we'll use sparse sets to lookup whether things duck type or not
	allMethods := map[canonicalMethod]int{}

	for _, iface := range interfaces {
		for method := range iface.Methods {
			if _, ok := allMethods[method]; !ok {
				allMethods[method] = len(allMethods)
			}
		}
	}

	for _, ty := range concreteTypes {
		for method := range ty.Methods {
			if _, ok := allMethods[method]; !ok {
				allMethods[method] = len(allMethods)
			}
		}
	}

	// Create a map of method names to corresponding interfaces
	interfaceToMethodSet := map[*scip.SymbolInformation]*intsets.Sparse{}
	for _, iface := range interfaces {
		if iface.Ident == nil {
			continue
		}

		methodSet := &intsets.Sparse{}
		for method := range iface.Methods {
			methodSet.Insert(allMethods[method])
		}

		interfaceToMethodSet[iface.Symbol] = methodSet
	}

	for _, ty := range concreteTypes {
		pos := ty.Ident.Pos()
		sym, ok := symbols.GetSymbolInformation(ty.Pkg, pos)
		if !ok {
			panic(fmt.Sprintf("Could not find symbol for %s", ty.Symbol))
		}

		methodSet := &intsets.Sparse{}
		for method := range ty.Methods {
			methodSet.Insert(allMethods[method])
		}

		tyImpls := implementationsForType(ty, methodSet, interfaceToMethodSet)
		for _, impl := range tyImpls {
			implDef, ok := interfaces[impl.Symbol]
			if !ok {
				fmt.Println(fmt.Sprintf("Could not find interface %s", impl.Symbol))
				continue
			}

			// Add implementation details for the struct & interface relationship
			sym.Relationships = append(sym.Relationships, &scip.Relationship{
				Symbol:           impl.Symbol,
				IsImplementation: true,
			})

			// For all methods, add imlementation details as well
			for name, implMethod := range implDef.Methods {
				tyMethodInfo, ok := ty.Methods[name]
				if !ok {
					continue
				}

				tyMethodInfo.Relationships = append(tyMethodInfo.Relationships, &scip.Relationship{
					Symbol:           implMethod.Symbol,
					IsImplementation: true,
				})
			}
		}
	}
}

func AddImplementationRelationships(pkgs loader.PackageLookup, allPackages loader.PackageLookup, symbols *lookup.Global) {
	output.WithProgress("Indexing Implementations", func() error {
		localInterfaces, localTypes, err := extractInterfacesAndConcreteTypes(pkgs, symbols)
		if err != nil {
			return err
		}

		remotePackages := make(loader.PackageLookup)
		for pkgID, pkg := range allPackages {
			if _, ok := pkgs[pkgID]; ok {
				continue
			}

			remotePackages[pkgID] = pkg
		}
		remoteInterfaces, _, err := extractInterfacesAndConcreteTypes(remotePackages, symbols)
		if err != nil {
			return err
		}

		// local type -> local interface
		findImplementations(localTypes, localInterfaces, symbols)

		// local type -> remote interface
		findImplementations(localTypes, remoteInterfaces, symbols)

		// TOOD: We should consider what this would even look like?
		//     I don't think this makes sense the current way that we are emitting
		//     implementations. You wouldn't even catch these anyways when uploading
		// remote type -> local interface
		// findImplementations(remoteTypes, localInterfaces, symbols)

		return nil
	})
}

func implementationsForType(ty ImplDef, tyMethods *intsets.Sparse, interfaceToMethodSet map[*scip.SymbolInformation]*intsets.Sparse) (matching []*scip.SymbolInformation) {
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
}

func extractInterfacesAndConcreteTypes(pkgs loader.PackageLookup, symbols *lookup.Global) (interfaces map[string]ImplDef, concreteTypes map[string]ImplDef, err error) {
	interfaces = map[string]ImplDef{}
	concreteTypes = map[string]ImplDef{}

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
				if obj.Exported() {
					handler.Println("No symbol for:", ident.Name, obj.Pkg(), obj.Id())
				}

				continue
			}

			methods := listMethods(objType)

			// ignore interfaces that are empty. they are too
			// plentiful and don't provide useful intelligence.
			if len(methods) == 0 {
				continue
			}

			canonicalizedMethods := map[canonicalMethod]*scip.SymbolInformation{}
			for _, m := range methods {
				// sym, ok := pkgSymbols.Get(m.Obj().Pos())
				sym, ok, err := symbols.GetSymbolOfObject(m.Obj())
				if err != nil {
					handler.ErrOrPanic("Error while looking for symbol %s | %s", err, m.Obj())
					continue
				}

				if !ok {
					// panic(fmt.Sprintf("Could not find symbol for %s", m.Obj()))
					continue
				}

				canonicalizedMethods[canonicalizeMethod(m)] = sym
			}

			d := ImplDef{
				Symbol:  symbol,
				Pkg:     pkg,
				Ident:   ident,
				Methods: canonicalizedMethods,
			}

			if types.IsInterface(objType) {
				interfaces[d.Symbol.Symbol] = d
			} else {
				concreteTypes[d.Symbol.Symbol] = d
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
func canonicalizeMethod(m *types.Selection) canonicalMethod {
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

	return canonicalMethod(builder.String())
}
