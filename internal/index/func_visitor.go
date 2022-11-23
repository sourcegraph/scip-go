package index

import (
	"go/ast"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitFunctionDefinition(doc *document.Document, pkg *packages.Package, node *ast.FuncDecl) {
	// Package
	desciptors := []*scip.Descriptor{
		{
			Name:   pkg.PkgPath,
			Suffix: scip.Descriptor_Namespace,
		},
	}

	// Receiver, if applicable
	if recv, has := receiverTypeName(node); has {
		desciptors = append(desciptors, descriptorType(recv))
	}

	// Name
	desciptors = append(desciptors, descriptorMethod(node.Name.Name))

	symbol := symbols.FromDescriptors(pkg, desciptors...)

	doc.SetNewSymbol(
		symbol,
		node,
		node.Name,
	)
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
