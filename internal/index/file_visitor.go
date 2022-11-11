package index

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"golang.org/x/tools/go/packages"
)

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

	// field definition position to symbol

	pkgSymbols    *lookup.Package
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
	case *ast.Ident:
		info := v.pkg.TypesInfo

		pos := node.NamePos
		position := v.pkg.Fset.Position(pos)

		// Emit Definition
		def := info.Defs[node]
		if def != nil {
			var sym string
			if pkgSymbols, ok := v.pkgSymbols.Get(def.Pos()); ok {
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
			var symbol string
			if localSymbol, ok := v.locals[ref.Pos()]; ok {
				symbol = localSymbol
			} else {
				refPkgPath := symbols.PkgPathFromObject(ref)
				pkg, ok := v.pkgLookup[refPkgPath]
				if !ok {
					panic(fmt.Sprintf("Failed to find the thing for ref: |%+v|\n", ref))
				}

				var err error
				symbol, ok, err = v.globalSymbols.GetSymbolOfObject(pkg, ref)
				if err != nil {
					fmt.Println("ERROR:", err)
					fmt.Println(pkg.Fset.Position(node.Pos()))
					return v
				}

				if !ok {
					return v
				}
			}

			v.doc.AppendSymbolReference(symbol, scipRange(position, ref))
		}
	}

	return v
}
