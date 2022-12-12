package visitors

import (
	"go/ast"
	"go/token"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitVarDefinition(doc *document.Document, pkg *packages.Package, decl *ast.GenDecl) {
	ast.Walk(varVisitor{
		doc:   doc,
		pkg:   pkg,
		scope: NewScope(pkg.PkgPath),
	}, decl)
}

type varVisitor struct {
	doc *document.Document
	pkg *packages.Package

	curDecl ast.Decl
	scope   *Scope
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
			v.doc.SetNewSymbol(symbol, v.curDecl, name)

			v.scope.push(name.Name, scip.Descriptor_Meta)
			walkExprList(v, node.Values)
			v.scope.pop()
		}

		return nil

	case *ast.Field:
		// I think the only case of this is embedded fields.
		if len(node.Names) == 0 {

			names := getIdentOfTypeExpr(v.pkg, node.Type)
			for _, name := range names {
				embeddedSymbol := v.makeSymbol(&scip.Descriptor{
					Name:   name.Name,
					Suffix: scip.Descriptor_Term,
				})

				// In this odd scenario, the definition is at embedded field level,
				// not wherever the name is. So that messes up our lookup table.
				v.doc.SetNewSymbolForPos(embeddedSymbol, node, name, node.Pos())
			}
		} else {
			for _, name := range node.Names {
				v.doc.SetNewSymbol(v.makeSymbol(&scip.Descriptor{
					Name:   name.Name,
					Suffix: scip.Descriptor_Term,
				}), nil, name)

				switch typ := node.Type.(type) {
				case *ast.MapType:
					v.scope.push(name.Name, scip.Descriptor_Term)
					defer func() {
						v.scope.pop()
					}()

					ast.Walk(v, typ.Key)
					ast.Walk(v, typ.Value)

				case *ast.ArrayType:
					v.scope.push(name.Name, scip.Descriptor_Term)
					defer func() {
						v.scope.pop()
					}()

					ast.Walk(v, typ.Elt)

				case *ast.StructType, *ast.InterfaceType:
					// Current scope is now embedded in the anonymous struct
					//   So we walk the rest of the type expression and save
					//   the nested names
					v.scope.push(name.Name, scip.Descriptor_Term)
					defer func() {
						v.scope.pop()
					}()

					ast.Walk(v, node.Type)
				}
			}
		}
		return nil

	default:
		return v
	}
}

func (s *varVisitor) makeSymbol(descriptor *scip.Descriptor) string {
	return symbols.FromDescriptors(s.pkg, append(s.scope.descriptors, descriptor)...)
}
