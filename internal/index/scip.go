package index

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func Index(opts config.IndexOpts) (*scip.Index, error) {
	return IndexProject(opts)
}

func IndexProject(opts config.IndexOpts) (*scip.Index, error) {
	opts.ModuleRoot, _ = filepath.Abs(opts.ModuleRoot)
	moduleRoot := opts.ModuleRoot

	pkgs, pkgLookup, err := loader.LoadPackages(opts, moduleRoot)
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
			ProjectRoot:          "file://" + moduleRoot,
			TextDocumentEncoding: scip.TextEncoding_UTF8,
		},
		Documents:       []*scip.Document{},
		ExternalSymbols: []*scip.SymbolInformation{},
	}

	pathToDocuments := map[string]*document.Document{}
	globalSymbols := lookup.NewGlobalSymbols()

	// We have to visit all the packages to get the definition sites
	// for all the symbols.
	//
	// We don't want to visit in the same depth as file visitors though,
	// so we do ONLY do this
	for _, pkg := range pkgLookup {
		visitPackage(moduleRoot, pkg, pathToDocuments, globalSymbols)
	}

	for _, pkg := range pkgs {
		pkgSymbols := globalSymbols.GetPackage(pkg)

		for _, f := range pkg.Syntax {
			relative, _ := filepath.Rel(moduleRoot, pkg.Fset.File(f.Package).Name())
			doc := pathToDocuments[relative]

			visitor := FileVisitor{
				doc:       doc,
				pkg:       pkg,
				file:      f,
				pkgLookup: pkgLookup,

				// locals are per-file, so create a new one per file
				locals: map[token.Pos]string{},

				pkgSymbols:    pkgSymbols,
				globalSymbols: globalSymbols,
			}

			// Generate import references
			for _, spec := range f.Imports {
				importedPackage := pkg.Imports[strings.Trim(spec.Path.Value, `"`)]
				if importedPackage == nil {
					fmt.Println("Could not find: ", spec.Path)
					continue
				}

				position := pkg.Fset.Position(spec.Pos())
				emitImportReference(doc, position, importedPackage)
			}

			ast.Walk(visitor, f)
			index.Documents = append(index.Documents, doc.Document)
		}
	}

	return &index, nil
}

func emitImportReference(
	doc *document.Document,
	position token.Position,
	importedPackage *packages.Package,
) {
	pkgPath := importedPackage.PkgPath
	scipRange := scipRangeFromName(position, pkgPath, true)
	symbol := symbols.FromDescriptors(importedPackage, descriptorPackage(pkgPath))

	doc.AppendSymbolReference(symbol, scipRange)
}

func scipRangeFromName(position token.Position, name string, adjust bool) []int32 {
	var adjustment int32 = 0
	if adjust {
		adjustment = 1
	}

	line := int32(position.Line - 1)
	column := int32(position.Column - 1)
	n := int32(len(name))

	return []int32{line, column + adjustment, column + n + adjustment}
}

func scipRange(position token.Position, obj types.Object) []int32 {
	var adjustment int32 = 0
	if pkgName, ok := obj.(*types.PkgName); ok && strings.HasPrefix(pkgName.Name(), `"`) {
		adjustment = 1
	}

	line := int32(position.Line - 1)
	column := int32(position.Column - 1)
	n := int32(len(obj.Name()))

	return []int32{line, column + adjustment, column + n - adjustment}
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

func visitPackage(
	moduleRoot string,
	pkg *packages.Package,
	pathToDocuments map[string]*document.Document,
	globalSymbols *lookup.Global,
) {
	pkgSymbols := lookup.NewPackageSymbols(pkg)

	// Iterate over all the files, collect any global symbols
	for _, f := range pkg.Syntax {
		relative, _ := filepath.Rel(moduleRoot, pkg.Fset.File(f.Package).Name())

		doc := visitSyntax(pkg, pkgSymbols, f, relative)

		// Save document for pass 2
		pathToDocuments[relative] = doc
	}

	globalSymbols.Add(pkgSymbols)
}

func visitSyntax(pkg *packages.Package, pkgSymbols *lookup.Package, f *ast.File, relative string) *document.Document {
	doc := document.NewDocument(relative, pkg, pkgSymbols)

	// TODO: Maybe we should do this before? we have traverse all
	// the fields first before, but now I think it's fine right here
	// .... maybe
	visitFieldsInFile(doc, pkg, f)

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.BadDecl:
			continue

		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				// These do not create global symbols
				continue

			case token.VAR, token.CONST:
				// visit var
				visitVarDefinition(doc, pkg, decl)

			case token.TYPE:
				// visitTypeDefinition(doc, pkg, decl)

			default:
				panic("Unhandled general declaration")
			}

		case *ast.FuncDecl:
			visitFunctionDefinition(doc, pkg, decl)
		}

	}

	return doc
}

func descriptorType(name string) *scip.Descriptor {
	return &scip.Descriptor{
		Name:   name,
		Suffix: scip.Descriptor_Type,
	}
}

func descriptorMethod(name string) *scip.Descriptor {
	return &scip.Descriptor{
		Name:   name,
		Suffix: scip.Descriptor_Method,
	}
}

func descriptorPackage(name string) *scip.Descriptor {
	return &scip.Descriptor{
		Name:   name,
		Suffix: scip.Descriptor_Package,
	}
}

func descriptorTerm(name string) *scip.Descriptor {
	return &scip.Descriptor{
		Name:   name,
		Suffix: scip.Descriptor_Term,
	}
}
