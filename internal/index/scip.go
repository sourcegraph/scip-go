package index

import (
	_ "embed"
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
var ScipGoVersion = strings.TrimSpace(versionFile)

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
	projectPackages, allPackages, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return nil, err
	}

	pathToDocuments := map[string]*document.Document{}
	for _, pkg := range allPackages {
		visitors.VisitPackageSyntax(
			opts.ModuleRoot, pkg, pathToDocuments, lookup.NewGlobalSymbols())
	}

	for _, pkg := range projectPackages {
		for _, f := range pkg.Syntax {
			docName := pkg.Fset.File(f.Package).Name()
			if _, ok := pathToDocuments[docName]; !ok {
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

	projectPackages, allPackages, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return err
	}

	pathToDocument, globalSymbols := indexVisitPackages(opts, projectPackages, allPackages)
	if !opts.SkipImplementations {
		if err := impls.AddImplementationRelationships(
			projectPackages, allPackages, globalSymbols,
		); err != nil {
			return err
		}
	}

	pkgIDs := slices.Sorted(maps.Keys(projectPackages))
	pkgLen := len(pkgIDs)

	var count uint64
	var wg sync.WaitGroup
	var writeErr error
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, ID := range pkgIDs {
			pkg := projectPackages[ID]
			pkgSymbols := globalSymbols.GetPackage(pkg)

			for _, file := range pkg.Syntax {
				doc := pathToDocument[pkg.Fset.File(file.Package).Name()]
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
					file,
					allPackages,
					pkgSymbols,
					globalSymbols,
				)

				// Traverse the file
				ast.Walk(visitor, file)

				// Write the document
				if writeErr = writer(visitor.ToScipDocument()); writeErr != nil {
					return
				}
			}

			atomic.AddUint64(&count, 1)
		}
	}()

	output.WithProgressParallel(&wg, "Visiting Project Files", &count, uint64(pkgLen))

	return writeErr
}

func indexVisitPackages(
	opts config.IndexOpts,
	projectPackages loader.PackageLookup,
	allPackages loader.PackageLookup,
) (map[string]*document.Document, *lookup.Global) {
	pathToDocuments := map[string]*document.Document{}
	globalSymbols := lookup.NewGlobalSymbols()

	var count uint64
	var wg sync.WaitGroup
	wg.Add(1)

	lookupIDs := slices.Sorted(maps.Keys(allPackages))

	// We have to visit all the packages to get the definition sites
	// for all the symbols.
	//
	// We don't want to visit in the same depth as file visitors though,
	// so we do ONLY do this
	go func() {
		defer wg.Done()

		for _, pkgID := range lookupIDs {
			pkg := allPackages[pkgID]
			slog.Debug("Visiting package", "path", pkg.PkgPath)
			visitors.VisitPackageSyntax(opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)

			pkgSymbol := globalSymbols.SetPkgSymbol(pkg)

			// If we don't have this package anywhere, don't try to create a new symbol
			if _, ok := projectPackages[newtypes.GetID(pkg)]; !ok {
				atomic.AddUint64(&count, 1)
				continue
			}

			symInfo := &scip.SymbolInformation{
				Symbol:        pkgSymbol,
				DisplayName:   pkg.Name,
				Documentation: findPackageDocs(pkg),
				SignatureDocumentation: &scip.Document{
					Language: "go",
					Text:     "package " + pkg.Name,
				},
			}
			firstFile := pkg.Syntax[0]
			firstDoc := pathToDocuments[pkg.Fset.File(firstFile.Package).Name()]
			firstDoc.SetSymbolInformation(firstFile.Name.NamePos, symInfo)

			for _, f := range pkg.Syntax {
				doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]
				position := pkg.Fset.Position(f.Name.NamePos)

				doc.PackageOccurrence = &scip.Occurrence{
					Range:       symbols.RangeFromName(position, f.Name.Name, false),
					Symbol:      pkgSymbol,
					SymbolRoles: int32(scip.SymbolRole_Definition),
				}
			}

			atomic.AddUint64(&count, 1)
		}
	}()

	output.WithProgressParallel(&wg, "Visiting Packages", &count, uint64(len(lookupIDs)))

	return pathToDocuments, globalSymbols
}
