package index

import (
	_ "embed"
	"fmt"
	"go/ast"
	"log/slog"
	"maps"
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/document"
	impls "github.com/sourcegraph/scip-go/internal/implementations"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/output"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip-go/internal/visitors"
	"google.golang.org/protobuf/proto"
)

//go:embed version.txt
var versionFile string
var ScipGoVersion string = strings.TrimSpace(versionFile)

type documentLookup = map[string]*document.Document

func GetPackages(opts config.IndexOpts) (current []newtypes.PackageID, deps []newtypes.PackageID, err error) {
	pkgs, pkgLookup, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return nil, nil, err
	}

	for name := range pkgs {
		current = append(current, name)
	}

	sort.Slice(current, func(i, j int) bool {
		return current[i] < current[j]
	})

	for name := range pkgLookup {
		deps = append(deps, name)
	}

	sort.Slice(deps, func(i, j int) bool {
		return deps[i] < deps[j]
	})

	return
}

func ListMissing(opts config.IndexOpts) (missing []string, err error) {
	pathToDocuments := map[string]*document.Document{}
	globalSymbols := lookup.NewGlobalSymbols()

	pkgs, pkgLookup, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return nil, err
	}

	lookupNames := slices.Sorted(maps.Keys(pkgLookup))
	for _, pkgName := range lookupNames {
		pkg := pkgLookup[pkgName]
		visitors.VisitPackageSyntax(opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)
	}

	pkgNames := slices.Sorted(maps.Keys(pkgs))
	for _, name := range pkgNames {
		pkg := pkgs[name]
		for _, f := range pkg.Syntax {
			docName := pkg.Fset.File(f.Package).Name()
			doc := pathToDocuments[docName]
			if doc == nil {
				missing = append(missing, docName)
			}

		}
	}

	return missing, nil
}

func Index(writer func(proto.Message) error, opts config.IndexOpts) error {
	// Emit Metadata.
	//   NOTE: Must be the first field emitted
	if err := writer(&scip.Metadata{
		Version: 0,
		ToolInfo: &scip.ToolInfo{
			Name:      "scip-go",
			Version:   ScipGoVersion,
			Arguments: []string{},
		},
		ProjectRoot:          "file://" + opts.ModuleRoot,
		TextDocumentEncoding: scip.TextEncoding_UTF8,
	}); err != nil {
		return err
	}

	pkgs, allPackages, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return err
	}

	pathToDocuments, globalSymbols := indexVisitPackages(opts, pkgs, allPackages)

	if !opts.SkipImplementations {
		impls.AddImplementationRelationships(pkgs, allPackages, globalSymbols)
	}

	pkgIDs := slices.Sorted(maps.Keys(pkgs))
	pkgLen := len(pkgIDs)

	var count uint64
	var wg sync.WaitGroup
	var writeErr error
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, ID := range pkgIDs {
			pkg := pkgs[ID]
			pkgSymbols := globalSymbols.GetPackage(pkg)

			for _, f := range pkg.Syntax {
				doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]
				if doc == nil {
					continue
				}

				// If possible, any state required for created a scip document
				// should be contained in the visitor. This makes sure that we can
				// garbage collect everything that's there after each loop,
				// rather than holding on to every occurrence and piece of data
				visitor := visitors.NewFileVisitor(
					doc,
					pkg,
					f,
					allPackages,
					pkgSymbols,
					globalSymbols,
				)

				// Traverse the file
				ast.Walk(visitor, f)

				// Write the document
				if writeErr = writer(visitor.ToScipDocument()); writeErr != nil {
					return
				}
			}

			atomic.AddUint64(&count, 1)
		}
	}()

	output.WithProgressParallel(&wg, "Visiting Project Files: ", &count, uint64(pkgLen))

	return writeErr
}

func indexVisitPackages(
	opts config.IndexOpts,
	pkgs loader.PackageLookup,
	pkgLookup loader.PackageLookup,
) (documentLookup, *lookup.Global) {
	pathToDocuments := documentLookup{}
	globalSymbols := lookup.NewGlobalSymbols()

	var count uint64
	var wg sync.WaitGroup
	wg.Add(1)

	lookupIDs := slices.Sorted(maps.Keys(pkgLookup))

	// We have to visit all the packages to get the definition sites
	// for all the symbols.
	//
	// We don't want to visit in the same depth as file visitors though,
	// so we do ONLY do this
	go func() {
		defer wg.Done()

		for _, pkgID := range lookupIDs {
			pkg := pkgLookup[pkgID]
			slog.Debug("Visiting package", "path", pkg.PkgPath)
			visitors.VisitPackageSyntax(opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)

			// Handle that packages can have many files for one package.
			// This finds the "definitive" package declaration
			pkgDeclaration, err := findBestPackageDefinitionPath(pkg)
			if err != nil {
				panic(fmt.Sprintf("Unhandled package declaration: %s", err))
			}

			if pkgDeclaration == nil {
				atomic.AddUint64(&count, 1)
				continue
			}

			globalSymbols.SetPkgName(pkg, pkgDeclaration)

			// If we don't have this package anywhere, don't try to create a new symbol
			if _, ok := pkgs[newtypes.GetID(pkg)]; !ok {
				atomic.AddUint64(&count, 1)
				continue
			}

			pkgSymbol := globalSymbols.GetPkgNameSymbol(pkg).Symbol
			for _, f := range pkg.Syntax {
				doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]

				if pkgDeclaration != nil {
					position := pkg.Fset.Position(f.Name.NamePos)

					role := int32(scip.SymbolRole_ReadAccess)
					if f == pkgDeclaration {
						doc.SetNewSymbolForPos(pkgSymbol, pkgDeclaration, f.Name, f.Name.NamePos)
						role = int32(scip.SymbolRole_Definition)
					}

					doc.PackageOccurrence = &scip.Occurrence{
						Range:       symbols.RangeFromName(position, f.Name.Name, false),
						Symbol:      pkgSymbol,
						SymbolRoles: role,
					}
				}
			}

			atomic.AddUint64(&count, 1)
		}
	}()

	output.WithProgressParallel(&wg, "Visiting Packages", &count, uint64(len(lookupIDs)))

	return pathToDocuments, globalSymbols
}

// packagePrefixes returns all prefix of the go package path. For example, the package
// `foo/bar/baz` will return the slice containing `foo/bar/baz`, `foo/bar`, and `foo`.
func packagePrefixes(packageName string) []string {
	parts := strings.Split(packageName, "/")
	prefixes := make([]string, len(parts))

	for i := 1; i <= len(parts); i++ {
		prefixes[len(parts)-i] = strings.Join(parts[:i], "/")
	}

	return prefixes
}
