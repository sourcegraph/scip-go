package visitors

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"golang.org/x/tools/go/packages"
)

func visitVarDefinition(doc *document.Document, pkg *packages.Package, decl *ast.GenDecl) {
	ast.Walk(varVisitor{
		doc: doc,
		pkg: pkg,
	}, decl)
}

type varVisitor struct {
	doc *document.Document
	pkg *packages.Package

	curDecl ast.Decl
}

var _ ast.Visitor = &varVisitor{}

func (v varVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.GenDecl:
		switch node.Tok {
		// Only traverse vars and consts
		case token.VAR, token.CONST:
			v.curDecl = node
			return v
		default:
			return nil
		}
	case *ast.ValueSpec:
		// Iterate over names, which are the only thing that can be definitions
		for _, name := range node.Names {
			symbol := symbols.FromDescriptors(v.pkg, descriptorTerm(name.Name))
			fmt.Println("symbol", symbol)
			v.doc.SetNewSymbol(symbol, v.curDecl, name)

			// position := v.pkg.Fset.Position(name.Pos())
			// v.doc.NewOccurrence(symbol, scipRangeFromName(position, name.Name, false))
		}

		for _, value := range node.Values {
			fmt.Printf("value: %s %T\n", value, value)
		}

		walkExprList(v, node.Values)

		return nil

	case *ast.CompositeLit:
		fmt.Println("composite lit", node)
		return v

	case *ast.StructType:
		// inline struct
		fmt.Println("struct", node)
		return v

	default:
		return nil
	}
}
