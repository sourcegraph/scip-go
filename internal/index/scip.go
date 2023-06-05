package index

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"

	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/funk"
	"github.com/sourcegraph/scip-go/internal/handler"
	impls "github.com/sourcegraph/scip-go/internal/implementations"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/output"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip-go/internal/visitors"
	"github.com/sourcegraph/scip/bindings/go/scip"
)

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

	lookupNames := funk.Keys(pkgLookup)
	for _, pkgName := range lookupNames {
		pkg := pkgLookup[pkgName]
		visitors.VisitPackageDeclarations(opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)
	}

	pkgNames := funk.Keys(pkgs)
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

func Index(opts config.IndexOpts) (*scip.Index, error) {
	pkgs, pkgLookup, err := loader.LoadPackages(opts, opts.ModuleRoot)
	if err != nil {
		return nil, err
	}

	index := scip.Index{
		Metadata: &scip.Metadata{
			Version: 0,
			ToolInfo: &scip.ToolInfo{
				Name:      "scip-go",
				Version:   "0.1",
				Arguments: []string{},
			},
			ProjectRoot:          "file://" + opts.ModuleRoot,
			TextDocumentEncoding: scip.TextEncoding_UTF8,
		},
		Documents:       []*scip.Document{},
		ExternalSymbols: []*scip.SymbolInformation{},
	}

	pathToDocuments := map[string]*document.Document{}
	globalSymbols := lookup.NewGlobalSymbols()

	output.WithProgress("Visiting Packages", func() error {
		// We have to visit all the packages to get the definition sites
		// for all the symbols.
		//
		// We don't want to visit in the same depth as file visitors though,
		// so we do ONLY do this
		lookupIDs := funk.Keys(pkgLookup)
		for _, pkgID := range lookupIDs {
			pkg := pkgLookup[pkgID]
			visitors.VisitPackageDeclarations(opts.ModuleRoot, pkg, pathToDocuments, globalSymbols)

			// TODO: I don't like this
			pkgDeclaration, err := findBestPackageDefinitionPath(pkg)
			if err != nil {
				panic(fmt.Sprintf("Unhandled package declaration: %s", err))
			}

			if pkgDeclaration == nil {
				continue
			}

			globalSymbols.SetPkgName(pkg, pkgDeclaration)

			if _, ok := pkgs[newtypes.GetID(pkg)]; !ok {
				continue
			}

			pkgSymbol := globalSymbols.GetPkgNameSymbol(pkg).Symbol
			for _, f := range pkg.Syntax {
				doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]

				if pkgDeclaration != nil {
					if f == pkgDeclaration {
						position := pkg.Fset.Position(f.Name.NamePos)
						doc.SetNewSymbolForPos(pkgSymbol, nil, f.Name, f.Name.NamePos)
						doc.NewDefinition(pkgSymbol, symbols.RangeFromName(position, f.Name.Name, false))
					} else {
						position := pkg.Fset.Position(f.Name.NamePos)
						doc.AppendSymbolReference(pkgSymbol, symbols.RangeFromName(position, f.Name.Name, false), nil)
					}
				}
			}

		}

		return nil
	})

	if opts.SkipImplementations {
		output.Println("Skipping implementation relationships\n")
	} else {
		impls.AddImplementationRelationships(pkgs, pkgLookup, globalSymbols)
	}

	output.WithProgress("Declaring Symbols for Documents", func() error {
		for _, doc := range pathToDocuments {
			doc.DeclareSymbols()
		}
		return nil
	})

	output.WithProgress("Visiting Project Files", func() error {
		pkgIDs := funk.Keys(pkgs)
		for _, ID := range pkgIDs {
			pkg := pkgs[ID]

			pkgSymbols := globalSymbols.GetPackage(pkg)

			for _, f := range pkg.Syntax {
				doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]
				if doc == nil {
					handler.Println("doc is nil for:", pkg.Fset.File(f.Package).Name())
					continue
				}

				visitor := visitors.NewFileVisitor(
					doc,
					pkg,
					f,
					pkgLookup,
					pkgSymbols,
					globalSymbols,
				)

				ast.Walk(visitor, f)
				index.Documents = append(index.Documents, doc.Document)
			}
		}
		return nil
	})

	return &index, nil
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
