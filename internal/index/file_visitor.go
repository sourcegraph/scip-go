package index

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"golang.org/x/tools/go/packages"
)

func NewFileVisitor(
	doc *document.Document,
	pkg *packages.Package,
	file *ast.File,
	pkgLookup map[string]*packages.Package,
	pkgSymbols *lookup.Package,
	globalSymbols *lookup.Global,
) *FileVisitor {
	caseClauses := map[token.Pos]types.Object{}
	for implicit, obj := range pkg.TypesInfo.Implicits {
		if _, ok := implicit.(*ast.CaseClause); ok {
			caseClauses[obj.Pos()] = obj
		}
	}

	return &FileVisitor{
		doc:           doc,
		pkg:           pkg,
		file:          file,
		pkgLookup:     pkgLookup,
		locals:        map[token.Pos]string{},
		caseClauses:   caseClauses,
		pkgSymbols:    pkgSymbols,
		globalSymbols: globalSymbols,
	}
}

// FileVisitor visits an entire file, but it must be called
// after StructVisitor.
//
// Iterates over a file,
type FileVisitor struct {
	// Document to append occurrences to
	doc *document.Document

	// Current file information
	pkg  *packages.Package
	file *ast.File

	// soething
	pkgLookup map[string]*packages.Package

	// local definition position to symbol
	locals map[token.Pos]string

	// case statement clauses
	caseClauses map[token.Pos]types.Object

	// field definition position to symbol for the package
	pkgSymbols *lookup.Package

	// field definition position to symbol for the entire compliation
	globalSymbols *lookup.Global
}

// Implements ast.Visitor
var _ ast.Visitor = &FileVisitor{}

func (f *FileVisitor) createNewLocalSymbol(pos token.Pos) string {
	if _, ok := f.locals[pos]; ok {
		panic("Cannot create a new local symbol for an ident that has already been created")
	}

	f.locals[pos] = fmt.Sprintf("local %d", len(f.locals))
	return f.locals[pos]
}

func (v FileVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.ImportSpec:
		// Skip imports
		return nil

	case *ast.SelectorExpr:
		if ident, ok := node.X.(*ast.Ident); ok {
			use := v.pkg.TypesInfo.Uses[ident]

			// We special case handling PkgNames because they do some goofy things
			// compared to almost every other construct in the language.
			switch sel := use.(type) {
			case *types.PkgName:
				pos := ident.NamePos
				position := v.pkg.Fset.Position(pos)
				symbol := v.globalSymbols.GetPkgNameSymbol(sel.Imported().Path()).Symbol
				v.doc.AppendSymbolReference(symbol, scipRange(position, sel), nil)

				// Then walk the selection
				ast.Walk(v, node.Sel)

				// and since we've handled the rest, end visit
				return nil
			}
		}

		return v
	case *ast.File:
		// ast.Walk(v, node.Name)
		// sym := v.globalSymbols.GetPkgNameSymbol(node.Pos())
		// fmt.Println("FILE", node.Name, sym)

		if node.Doc != nil {
			ast.Walk(v, node.Doc)
		}

		// TODO: Handle package name declaration separately
		// ast.Walk(v, n.Name)

		walkDeclList(v, node.Decls)
		return nil
	case *ast.Ident:
		// Short circuit if this is a blank identifier
		if node.Name == "_" {
			return nil
		}

		pos := node.NamePos
		position := v.pkg.Fset.Position(pos)

		// Short circuit on case clauses
		if obj, ok := v.caseClauses[node.Pos()]; ok {
			sym := v.createNewLocalSymbol(obj.Pos())
			v.doc.NewOccurrence(sym, scipRange(position, obj))
			return nil
		}

		info := v.pkg.TypesInfo

		// Emit Definition
		def := info.Defs[node]
		if def != nil {
			var sym string
			if pkgSymbols, ok := v.pkgSymbols.GetSymbol(def.Pos()); ok {
				sym = pkgSymbols
			} else if globalSymbol, ok := v.globalSymbols.GetSymbol(v.pkg, def.Pos()); ok {
				sym = globalSymbol
			} else {
				sym = v.createNewLocalSymbol(def.Pos())
			}

			v.doc.NewOccurrence(sym, scipRange(position, def))
		}

		// Emit Reference
		ref := info.Uses[node]
		if ref != nil {
			var (
				symbol       string
				overrideType types.Type
			)

			if localSymbol, ok := v.locals[ref.Pos()]; ok {
				symbol = localSymbol

				if _, ok := v.caseClauses[ref.Pos()]; ok {
					overrideType = v.pkg.TypesInfo.TypeOf(node)
				}
			} else {
				var err error
				symInfo, ok, err := v.globalSymbols.GetSymbolOfObject(ref)
				if err != nil {
					_, ok := v.pkgLookup[symbols.PkgPathFromObject(ref)]
					if !ok {
						panic(fmt.Sprintf("Failed to find a package for ref: |%+v|\n", ref))
					}

					panic(fmt.Sprintf("Unable to find symbol of object: %s", err))
				}

				if !ok {
					return v
				}

				// Set the resulting info
				symbol = symInfo.Symbol
			}

			v.doc.AppendSymbolReference(symbol, scipRange(position, ref), overrideType)
		}

		if def == nil && ref == nil {
			panic(fmt.Sprintf(
				"Neither def nor ref found: %s | %T | %s",
				node.Name,
				node,
				v.pkg.Fset.Position(node.Pos())),
			)
		}
	}

	return v
}

func walkDeclList(v ast.Visitor, list []ast.Decl) {
	for _, x := range list {
		ast.Walk(v, x)
	}
}
