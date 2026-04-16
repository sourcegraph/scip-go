package impls

import (
	"fmt"
	"go/types"
	"log/slog"

	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/output"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

// methodID is a unique identifier for a method, using types.Id semantics
// (package-path-qualified for unexported methods, just the name for exported).
type methodID string

type ImplDef struct {
	// The corresponding scip symbol
	Symbol *scip.SymbolInformation

	Pkg     *packages.Package
	Named   *types.Named
	Methods map[methodID]*scip.SymbolInformation
}

func findImplementations(concreteTypes map[string]ImplDef, interfaces map[string]ImplDef, symbols *lookup.Global) {
	for _, ty := range concreteTypes {
		for _, iface := range interfaces {
			ifaceType, ok := iface.Named.Underlying().(*types.Interface)
			if !ok {
				continue
			}

			// Use types.Implements to correctly check interface satisfaction,
			// handling all edge cases (embedded types, generics, unexported methods).
			if !types.Implements(ty.Named, ifaceType) &&
				!types.Implements(types.NewPointer(ty.Named), ifaceType) {
				continue
			}

			// Add implementation details for the struct & interface relationship
			ty.Symbol.Relationships = append(ty.Symbol.Relationships,
				&scip.Relationship{
					Symbol:           iface.Symbol.Symbol,
					IsImplementation: true,
				})

			// For all methods, add implementation details as well
			for name, implMethod := range iface.Methods {
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

func AddImplementationRelationships(
	pkgs loader.PackageLookup,
	allPackages loader.PackageLookup,
	symbols *lookup.Global,
) ([]*scip.SymbolInformation, error) {
	var externalSymbols []*scip.SymbolInformation
	err := output.WithProgress("Indexing Implementations", func() error {
		var msCache typeutil.MethodSetCache
		localInterfaces, localTypes := extractInterfacesAndConcreteTypes(
			pkgs, symbols, &msCache)

		remotePackages := make(loader.PackageLookup)
		for pkgID, pkg := range allPackages {
			if _, ok := pkgs[pkgID]; ok {
				continue
			}

			remotePackages[pkgID] = pkg
		}
		remoteInterfaces, remoteTypes := extractInterfacesAndConcreteTypes(
			remotePackages, symbols, &msCache)

		// local type -> local interface
		findImplementations(localTypes, localInterfaces, symbols)

		// local type -> remote interface
		findImplementations(localTypes, remoteInterfaces, symbols)

		// remote type -> local interface
		// We emit these as external symbols so index consumer can merge them.
		findImplementations(remoteTypes, localInterfaces, symbols)

		// Collect remote type symbols that gained relationships
		for _, typ := range remoteTypes {
			if len(typ.Symbol.Relationships) > 0 {
				externalSymbols = append(externalSymbols, typ.Symbol)
			}
		}

		return nil
	})
	return externalSymbols, err
}

func extractInterfacesAndConcreteTypes(
	pkgs loader.PackageLookup,
	globalSymbols *lookup.Global,
	msCache *typeutil.MethodSetCache,
) (map[string]ImplDef, map[string]ImplDef) {
	interfaces := map[string]ImplDef{}
	concreteTypes := map[string]ImplDef{}

	for _, pkg := range pkgs {
		// Builtin isn't the same as standard library, that is for builtin types
		// We don't need to check those for implemenations.
		if pkg.Name == "builtin" {
			continue
		}

		if pkg.Types == nil {
			slog.Warn("No types for package", "path", pkg.PkgPath)
			continue
		}

		if pkg.TypesInfo != nil {
			extractFromTypesInfo(pkg, globalSymbols, msCache, interfaces, concreteTypes)
			continue
		}

		extractFromScope(pkg, globalSymbols, msCache, interfaces, concreteTypes)
	}

	return interfaces, concreteTypes
}

// extractFromTypesInfo handles project packages that have full TypesInfo/AST.
func extractFromTypesInfo(
	pkg *packages.Package,
	globalSymbols *lookup.Global,
	msCache *typeutil.MethodSetCache,
	interfaces map[string]ImplDef,
	concreteTypes map[string]ImplDef,
) {
	pkgSymbols := globalSymbols.GetPackage(pkg)
	if pkgSymbols == nil {
		slog.Warn("No symbols for package", "path", pkg.PkgPath)
		return
	}

	for ident, obj := range pkg.TypesInfo.Defs {
		if obj == nil {
			continue
		}

		obj, ok := obj.(*types.TypeName)
		if !ok {
			continue
		}

		// Skip types declared inside function bodies
		if pkg.Types != nil && obj.Parent() != pkg.Types.Scope() {
			continue
		}

		objType, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}

		symbol, ok := pkgSymbols.Get(obj.Pos())
		if !ok {
			slog.Debug(
				"No symbol for package-level named type",
				"identifier", ident.Name, "package", pkg.PkgPath, "id", obj.Id())
			continue
		}

		collectImplDef(
			pkg, globalSymbols, msCache, objType, symbol, interfaces, concreteTypes)
	}
}

// extractFromScope handles dep packages that have no AST — uses
// pkg.Types.Scope() to enumerate package-level named types and synthesizes
// their SCIP symbols.
func extractFromScope(
	pkg *packages.Package,
	globalSymbols *lookup.Global,
	msCache *typeutil.MethodSetCache,
	interfaces map[string]ImplDef,
	concreteTypes map[string]ImplDef,
) {
	scope := pkg.Types.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		tn, ok := obj.(*types.TypeName)
		if !ok {
			continue
		}

		objType, ok := tn.Type().(*types.Named)
		if !ok {
			continue
		}

		sym, ok := symbols.SynthesizeFromObject(pkg, tn)
		if !ok {
			slog.Debug("Could not synthesize symbol for dep type",
				"name", name, "package", pkg.PkgPath)
			continue
		}
		symbol := &scip.SymbolInformation{Symbol: sym}

		collectImplDef(pkg, globalSymbols, msCache, objType, symbol, interfaces, concreteTypes)
	}
}

func collectImplDef(
	pkg *packages.Package,
	globalSymbols *lookup.Global,
	msCache *typeutil.MethodSetCache,
	objType *types.Named,
	symbol *scip.SymbolInformation,
	interfaces map[string]ImplDef,
	concreteTypes map[string]ImplDef,
) {
	methods := typeutil.IntuitiveMethodSet(objType, msCache)

	// ignore interfaces that are empty. they are too
	// plentiful and don't provide useful intelligence.
	if len(methods) == 0 {
		return
	}

	methodIds := map[methodID]*scip.SymbolInformation{}
	for _, m := range methods {
		sym, ok, err := globalSymbols.GetSymbolOfObject(m.Obj())
		if err != nil {
			slog.Debug(fmt.Sprintf("Error while looking for symbol %s | %s", err, m.Obj()))
			continue
		}

		if !ok {
			continue
		}

		methodIds[methodID(m.Obj().Id())] = sym
	}

	d := ImplDef{
		Symbol:  symbol,
		Pkg:     pkg,
		Named:   objType,
		Methods: methodIds,
	}

	if types.IsInterface(objType) {
		interfaces[d.Symbol.Symbol] = d
	} else {
		concreteTypes[d.Symbol.Symbol] = d
	}
}
