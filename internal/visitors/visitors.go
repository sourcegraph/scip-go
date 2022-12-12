package visitors

import (
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func VisitPackageDeclarations(
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

func walkExprList(v ast.Visitor, list []ast.Expr) {
	for _, x := range list {
		ast.Walk(v, x)
	}
}

func walkDeclList(v ast.Visitor, list []ast.Decl) {
	for _, x := range list {
		ast.Walk(v, x)
	}
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
