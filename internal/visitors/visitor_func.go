package visitors

import (
	"go/ast"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

var _ ast.Visitor = &funcVisitor{}

type funcVisitor struct {
	pkg   *packages.Package
	doc   *document.Document
	scope *Scope
}

func (v funcVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.FuncDecl:
		// Receiver, if applicable
		if recv, has := receiverTypeName(node); has {
			v.scope.push(recv, scip.Descriptor_Type)
		}

		symbol := v.scope.makeSymbol(v.pkg, node.Name.Name, scip.Descriptor_Method)
		v.doc.SetNewSymbol(symbol, node, node.Name)

		// Any associated declarations should be generated with the scope of this method
		v.scope.push("func", scip.Descriptor_Meta)
		v.scope.push(node.Name.Name, scip.Descriptor_Meta)
		ast.Walk(v, node.Type)
		v.scope.pop()

		return nil

	case *ast.FuncType:
		// Should not need to declare any non-local definitions in the type params
		// if node.TypeParams != nil {
		// 	Walk(v, node.TypeParams)
		// }

		// Should not need to declare any non-local definitions in the params
		// if node.Params != nil {
		// 	Walk(v, node.Params)
		// }

		// Types can create new interfaces and/or types,
		// so we need to visit them and potentially declare new non-local symbols
		if node.Results != nil {
			ast.Walk(v, node.Results)
		}

		return nil

	case *ast.BlockStmt:
		return nil

	case *ast.InterfaceType:
		// TODO: Should handle this more elegantly?
		for _, field := range node.Methods.List {
			for _, name := range field.Names {
				symbol := v.scope.makeSymbol(v.pkg, name.Name, scip.Descriptor_Method)
				v.doc.SetNewSymbol(symbol, name, name)
			}
		}

		return nil

	default:
		return v
	}

}

func visitFunctionDefinition(doc *document.Document, pkg *packages.Package, node *ast.FuncDecl) {
	visitor := funcVisitor{
		pkg:   pkg,
		doc:   doc,
		scope: NewScope(pkg.PkgPath),
	}

	ast.Walk(visitor, node)
}

func receiverTypeName(f *ast.FuncDecl) (string, bool) {
	recv := f.Recv
	if recv == nil {
		return "", false
	}

	if len(recv.List) > 1 {
		panic("I don't understand what this would look like")
	} else if len(recv.List) == 0 {
		return "", false
	}

	field := recv.List[0]
	if field.Type == nil {
		return "", false
	}

	// Dereference pointer receiver types
	typ := field.Type
	if p, _ := typ.(*ast.StarExpr); p != nil {
		typ = p.X
	}

	// If we have an identifier, then we have a receiver
	if p, _ := typ.(*ast.Ident); p != nil {
		return p.Name, true
	}

	return "", false
}
