package visitors

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

var inlineCount = 0

func visitVarDefinition(doc *document.Document, pkg *packages.Package, decl *ast.GenDecl) {
	ast.Walk(&varVisitor{
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

func (v *varVisitor) Visit(n ast.Node) (w ast.Visitor) {
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
			symbol := v.makeSymbol(descriptorTerm(name.Name))
			v.doc.SetNewSymbol(symbol, v.curDecl, name)
		}

		if len(node.Names) > 0 {
			var scopeName string
			if len(node.Names) == 1 {
				scopeName = node.Names[0].Name
			} else {
				scopeName = fmt.Sprintf("inline-%d", inlineCount)
				inlineCount += 1

			}

			v.scope.push(scopeName, scip.Descriptor_Meta)
			if node.Type != nil {
				switch valueType := node.Type.(type) {
				case *ast.FuncType:
					// nothing
				case *ast.Ident:
					// nothing
				case *ast.ChanType:
					// TODO: Could have chan with new struct
				case *ast.ArrayType:
					// TODO: Could have array with new struct
				case *ast.MapType:
					// TODO: could have nested struct?
				case *ast.StarExpr:
					// TODO: could be *struct?
				case *ast.SelectorExpr:
					// TODO?
				case *ast.IndexExpr:
					// TODO: Generics, possibly need to travrerse
				case *ast.InterfaceType:
					ast.Walk(v, valueType)
				case *ast.StructType:
					// panic(fmt.Sprintf("TODO: handle type %T %s", valueType, v.pkg.Fset.Position(node.Pos())))
					ast.Walk(v, valueType)
				default:
					// TODO: Consider how we could emit errors for users running this, not in dev mode
					_ = handler.ErrOrPanic("TODO: handle type %T %s", valueType, v.pkg.Fset.Position(node.Pos()))

					return
				}
			}

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
