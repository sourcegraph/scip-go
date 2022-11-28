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
	impls "github.com/sourcegraph/scip-go/internal/implementations"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func Index(opts config.IndexOpts) (*scip.Index, error) {
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
		fmt.Println("Attempting pkg:", pkg.PkgPath)

		visitPackage(moduleRoot, pkg, pathToDocuments, globalSymbols)

		// TODO: I don't like this
		pkgDeclaration, err := findBestPackageDefinitionPath(pkg)
		if err != nil {
			panic(fmt.Sprintf("Unhandled package declaration: %s", err))
		}

		if pkgDeclaration == nil {
			continue
		}

		globalSymbols.SetPkgName(pkg, pkgDeclaration)

		if _, ok := pkgs[pkg.PkgPath]; !ok {
			continue
		}

		pkgSymbol := globalSymbols.GetPkgNameSymbol(pkg.PkgPath).Symbol.Symbol

		for _, f := range pkg.Syntax {
			doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]

			if pkgDeclaration != nil {
				if f == pkgDeclaration {
					fmt.Println("FOUND ONE FOR THIS PACKAGE", f.Name)
					position := pkg.Fset.Position(f.Name.NamePos)
					doc.SetNewSymbolForPos(pkgSymbol, nil, f.Name, f.Name.NamePos)
					doc.NewDefinition(pkgSymbol, scipRangeFromName(position, f.Name.Name, false))
				} else {
					position := pkg.Fset.Position(f.Name.NamePos)
					fmt.Println("Emit symbol:", pkgSymbol, position)
					doc.AppendSymbolReference(pkgSymbol, scipRangeFromName(position, f.Name.Name, false), nil)
				}
			}
		}

	}

	impls.AddImplementationRelationships(pkgs, globalSymbols)

	// NOTE:
	// I'm not sure how to do this yet... but we basically need to iterate over
	// all the possible implementations and other relationships. After doing so
	// is when we can add the symbols itself to the documents. It seems a bit weird
	// but I'll see if there's some other way to do it later.
	for _, doc := range pathToDocuments {
		doc.DeclareSymbols()
	}

	for _, pkg := range pkgs {
		pkgSymbols := globalSymbols.GetPackage(pkg)

		for _, f := range pkg.Syntax {
			doc := pathToDocuments[pkg.Fset.File(f.Package).Name()]

			visitor := NewFileVisitor(
				doc,
				pkg,
				f,
				pkgLookup,
				pkgSymbols,
				globalSymbols,
			)

			// Generate import references
			for _, spec := range f.Imports {
				importedPackage := pkg.Imports[strings.Trim(spec.Path.Value, `"`)]
				if importedPackage == nil {
					fmt.Println("Could not find: ", spec.Path)
					continue
				}

				position := pkg.Fset.Position(spec.Pos())
				emitImportReference(globalSymbols, doc, position, importedPackage)
			}

			ast.Walk(visitor, f)
			index.Documents = append(index.Documents, doc.Document)
		}
	}

	return &index, nil
}

func emitImportReference(
	globalSymbols *lookup.Global,
	doc *document.Document,
	position token.Position,
	importedPackage *packages.Package,
) {
	pkgPath := importedPackage.PkgPath
	scipRange := scipRangeFromName(position, pkgPath, true)
	symbol := globalSymbols.GetPkgNameSymbol(pkgPath)
	doc.AppendSymbolReference(symbol.Symbol.Symbol, scipRange, nil)
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

		abs := pkg.Fset.File(f.Package).Name()
		relative, _ := filepath.Rel(moduleRoot, abs)

		doc := visitSyntax(pkg, pkgSymbols, f, relative)

		// Save document for pass 2
		pathToDocuments[abs] = doc
	}

	globalSymbols.Add(pkgSymbols)
}

func visitSyntax(pkg *packages.Package, pkgSymbols *lookup.Package, f *ast.File, relative string) *document.Document {
	doc := document.NewDocument(relative, pkg, pkgSymbols)

	// TODO: Maybe we should do this before? we have traverse all
	// the fields first before, but now I think it's fine right here
	// .... maybe
	visitTypesInFile(doc, pkg, f)

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.BadDecl:
			continue

		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				// These do not create global symbols
				continue

			case token.TYPE:
				// We do this via visitTypesInFile above

			case token.VAR, token.CONST:
				// visit var
				visitVarDefinition(doc, pkg, decl)

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
