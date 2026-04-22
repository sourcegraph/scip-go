package impls

import (
	"fmt"
	"go/ast"
	"go/types"
	"hash/crc32"
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/scip-code/scip-go/internal/implementations/fingerprint"
	"github.com/scip-code/scip-go/internal/loader"
	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip-go/internal/output"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

// methodID is a unique identifier for a method, using types.Id semantics
// (package-path-qualified for unexported methods, just the name for exported).
type methodID string

type ImplDef struct {
	// The corresponding scip symbol, generated via previous iteration over the AST
	Symbol *scip.SymbolInformation

	Pkg     *packages.Package
	Ident   *ast.Ident
	Named   *types.Named
	Methods map[methodID]*scip.SymbolInformation

	// Precomputed fields for fast filtering before calling types.Implements.
	Mask           uint64 // bitmask of hashed method signatures
	MethodCount    int
	HasUnexported  bool   // true if any method is unexported
	PkgPath        string // package path, used to skip cross-package checks for unexported methods
}

// methodMask computes a 64-bit bitmask from a set of methods, using
// the same algorithm as gopls's methodsets package. Each method's
// identity (Id + fingerprint) is hashed and mapped to a single bit.
// The resulting mask enables O(1) rejection of non-matching
// type/interface pairs: if ty.Mask & iface.Mask != iface.Mask,
// the type is missing at least one method of the interface.
func methodMask(methods []*types.Selection) uint64 {
	var mask uint64
	var buf []byte
	for _, m := range methods {
		fn := m.Obj().(*types.Func)
		id := fn.Id()
		fp, _ := fingerprint.Encode(fn.Signature())
		buf = append(buf[:0], id...)
		buf = append(buf, fp...)
		sum := crc32.ChecksumIEEE(buf)
		mask |= 1 << uint64(((sum>>24)^(sum>>16)^(sum>>8)^sum)&0x3f)
	}
	return mask
}

// hasUnexportedMethods returns true if any of the given methods is unexported.
func hasUnexportedMethods(methods []*types.Selection) bool {
	for _, m := range methods {
		if !m.Obj().Exported() {
			return true
		}
	}
	return false
}

func findImplementations(concreteTypes map[string]ImplDef, interfaces map[string]ImplDef, symbols *lookup.Global, count *uint64) {
	for _, ty := range concreteTypes {
		pos := ty.Ident.Pos()
		sym, ok := symbols.GetSymbolInformation(ty.Pkg, pos)
		if !ok {
			panic(fmt.Sprintf("Could not find symbol for %s", ty.Symbol))
		}

		for _, iface := range interfaces {
			if iface.Ident == nil {
				continue
			}

			ifaceType, ok := iface.Named.Underlying().(*types.Interface)
			if !ok {
				continue
			}

			// Fast filter 1: if the interface has unexported methods,
			// only types in the same package can implement it.
			if iface.HasUnexported && ty.PkgPath != iface.PkgPath {
				continue
			}

			// Fast filter 2: the type must have at least as many methods
			// as the interface requires.
			if ty.MethodCount < iface.MethodCount {
				continue
			}

			// Fast filter 3: bitmask subset check. If the type's mask
			// doesn't cover all bits of the interface's mask, the type
			// is definitely missing at least one required method.
			if ty.Mask&iface.Mask != iface.Mask {
				continue
			}

			// Use types.Implements to correctly check interface satisfaction,
			// handling all edge cases (embedded types, generics, unexported methods).
			if !types.Implements(ty.Named, ifaceType) &&
				!types.Implements(types.NewPointer(ty.Named), ifaceType) {
				continue
			}

			// Add implementation details for the struct & interface relationship
			sym.Relationships = append(sym.Relationships, &scip.Relationship{
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

		atomic.AddUint64(count, 1)
	}
}

func AddImplementationRelationships(
	pkgs loader.PackageLookup,
	allPackages loader.PackageLookup,
	symbols *lookup.Global,
) ([]*scip.SymbolInformation, error) {
	var externalSymbols []*scip.SymbolInformation

	var msCache typeutil.MethodSetCache
	localInterfaces, localTypes, err := extractInterfacesAndConcreteTypes(
		pkgs, symbols, &msCache)
	if err != nil {
		return nil, err
	}

	remotePackages := make(loader.PackageLookup)
	for pkgID, pkg := range allPackages {
		if _, ok := pkgs[pkgID]; ok {
			continue
		}

		remotePackages[pkgID] = pkg
	}
	remoteInterfaces, remoteTypes, err := extractInterfacesAndConcreteTypes(
		remotePackages, symbols, &msCache)
	if err != nil {
		return nil, err
	}

	// Total concrete types to check across the three passes.
	total := uint64(len(localTypes)*2 + len(remoteTypes))
	var count uint64
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// local type -> local interface
		findImplementations(localTypes, localInterfaces, symbols, &count)

		// local type -> remote interface
		findImplementations(localTypes, remoteInterfaces, symbols, &count)

		// remote type -> local interface
		// We emit these as external symbols so index consumer can merge them.
		findImplementations(remoteTypes, localInterfaces, symbols, &count)
	}()

	output.WithProgressParallel(&wg, "Indexing Implementations", &count, total)

	// Collect remote type symbols that gained relationships
	for _, typ := range remoteTypes {
		if sym, ok := symbols.GetSymbolInformation(typ.Pkg, typ.Ident.Pos()); ok {
			if len(sym.Relationships) > 0 {
				externalSymbols = append(externalSymbols, sym)
			}
		}
	}

	return externalSymbols, nil
}

func extractInterfacesAndConcreteTypes(
	pkgs loader.PackageLookup,
	symbols *lookup.Global,
	msCache *typeutil.MethodSetCache,
) (interfaces map[string]ImplDef, concreteTypes map[string]ImplDef, err error) {
	interfaces = map[string]ImplDef{}
	concreteTypes = map[string]ImplDef{}

	for _, pkg := range pkgs {
		// Builtin isn't the same as standard library, that is for builtin types
		// We don't need to check those for implemenations.
		if pkg.Name == "builtin" {
			continue
		}

		if pkg.TypesInfo == nil {
			slog.Warn("No types for package", "path", pkg.PkgPath)
			continue
		}

		pkgSymbols := symbols.GetPackage(pkg)
		if pkgSymbols == nil {
			slog.Warn("No symbols for package", "path", pkg.PkgPath)
			continue
		}

		for ident, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}

			// We ignore aliases 'type M = N' to avoid duplicate reporting
			// of the Named type N.
			obj, ok := obj.(*types.TypeName)
			if !ok {
				continue
			}

			// Skip types declared inside function bodies — the type visitor
			// only indexes package-level declarations, so local types will
			// never have a symbol entry.
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

			methods := typeutil.IntuitiveMethodSet(objType, msCache)

			// ignore interfaces that are empty. they are too
			// plentiful and don't provide useful intelligence.
			if len(methods) == 0 {
				continue
			}

			methodIds := map[methodID]*scip.SymbolInformation{}
			for _, m := range methods {
				sym, ok, err := symbols.GetSymbolOfObject(m.Obj())
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
				Symbol:        symbol,
				Pkg:           pkg,
				Ident:         ident,
				Named:         objType,
				Methods:       methodIds,
				Mask:          methodMask(methods),
				MethodCount:   len(methods),
				HasUnexported: hasUnexportedMethods(methods),
				PkgPath:       pkg.PkgPath,
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
