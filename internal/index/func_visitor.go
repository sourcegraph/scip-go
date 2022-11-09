package index

import (
	"go/ast"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitFunctionDefinition(doc *document.Document, pkg *packages.Package, node *ast.FuncDecl) {
	desciptors := []*scip.Descriptor{
		{
			Name:   pkg.PkgPath,
			Suffix: scip.Descriptor_Namespace,
		},
	}
	if recv, has := receiverTypeName(node); has {
		desciptors = append(desciptors, descriptorType(recv))
	}
	desciptors = append(desciptors, descriptorMethod(node.Name.Name))
	symbol := scipSymbolFromDescriptors(pkg, desciptors)

	doc.DeclareNewSymbol(
		symbol,
		node,
		node.Name,
	)
}
