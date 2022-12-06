package index

import (
	"fmt"
	"go/ast"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitTypesInFile(doc *document.Document, pkg *packages.Package, file *ast.File) {
	visitor := TypeVisitor{
		pkg: pkg,
		doc: doc,
		curScope: []*scip.Descriptor{
			{
				Name:   pkg.PkgPath,
				Suffix: scip.Descriptor_Namespace,
			},
		},
	}

	ast.Walk(visitor, file)
}

// TypeVisitor collects the all the information for top-level structs
// that can be imported by any other file (they do not have to be exported).
//
// For example, a struct `myStruct` can be imported by other files in the same
// packages. So we need to make those field names global (we only have global
// or file-local).
type TypeVisitor struct {
	doc *document.Document
	pkg *packages.Package

	curScope    []*scip.Descriptor
	curDecl     *ast.GenDecl
	isInterface bool
}

func (v TypeVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.GenDecl:
		// Current declaration is required for some documentation parsing.
		// So we have to keep this here with us as we traverse more deeply
		v.curDecl = node

		if node.Doc != nil {
			ast.Walk(v, node.Doc)
		}

		for _, s := range node.Specs {
			switch spec := s.(type) {
			case *ast.TypeSpec:
				switch spec.Type.(type) {
				case *ast.InterfaceType:
					v.isInterface = true
				default:
					v.isInterface = false
				}
			}

			ast.Walk(v, s)
		}

		return nil

	case
		// Continue down file and decls
		*ast.File,

		// Toplevel types that are important
		*ast.StructType,
		*ast.InterfaceType,

		// Continue traversing subtypes
		*ast.FieldList,
		*ast.Ident:

		return v

	case *ast.TypeSpec:
		v.curScope = append(v.curScope, &scip.Descriptor{
			Name:   node.Name.Name,
			Suffix: scip.Descriptor_Type,
		})
		defer func() {
			v.curScope = v.curScope[:len(v.curScope)-1]
		}()

		v.doc.SetNewSymbol(
			symbols.FromDescriptors(v.pkg, v.curScope...),
			v.curDecl,
			node.Name,
		)

		ast.Walk(v, node.Type)
	case *ast.Field:
		// I think the only case of this is embedded fields.
		if len(node.Names) == 0 {
			// If we have an interface, these do not *declare* a new symbol,
			// they simply add another constraint.
			if v.isInterface {
				return v
			}

			names := v.getIdentOfTypeExpr(node.Type)
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
					v.curScope = append(v.curScope, &scip.Descriptor{
						Name:   name.Name,
						Suffix: scip.Descriptor_Term,
					})
					defer func() {
						v.curScope = v.curScope[:len(v.curScope)-1]
					}()

					ast.Walk(v, typ.Key)
					ast.Walk(v, typ.Value)

				case *ast.ArrayType:
					v.curScope = append(v.curScope, &scip.Descriptor{
						Name:   name.Name,
						Suffix: scip.Descriptor_Term,
					})
					defer func() {
						v.curScope = v.curScope[:len(v.curScope)-1]
					}()

					ast.Walk(v, typ.Elt)

				case *ast.StructType, *ast.InterfaceType:
					// Current scope is now embedded in the anonymous struct
					//   So we walk the rest of the type expression and save
					//   the nested names
					v.curScope = append(v.curScope, &scip.Descriptor{
						Name:   name.Name,
						Suffix: scip.Descriptor_Term,
					})
					defer func() {
						v.curScope = v.curScope[:len(v.curScope)-1]
					}()

					ast.Walk(v, node.Type)
				}
			}
		}
	}

	return nil
}

// Implements ast.Visitor
var _ ast.Visitor = &TypeVisitor{}

func (s *TypeVisitor) getIdentOfTypeExpr(ty ast.Expr) []*ast.Ident {
	switch ty := ty.(type) {
	case *ast.Ident:
		return []*ast.Ident{ty}
	case *ast.SelectorExpr:
		return []*ast.Ident{ty.Sel}
	case *ast.StarExpr:
		return s.getIdentOfTypeExpr(ty.X)
	case *ast.IndexExpr:
		return s.getIdentOfTypeExpr(ty.X)
	case *ast.BinaryExpr:
		// As far as I can tell, binary exprs are ONLY for type constraints
		// and those don't really define anything on the struct.
		//
		// So far now, we'll just not return anything.
		//
		// return append(s.getIdentOfTypeExpr(ty.X), s.getIdentOfTypeExpr(ty.Y)...)
		return []*ast.Ident{}
	case *ast.UnaryExpr:
		return s.getIdentOfTypeExpr(ty.X)
	default:
		panic(fmt.Sprintf("Unhandled named struct field: %T %+v\n%s", ty, ty, s.pkg.Fset.Position(ty.Pos())))
	}
}

func (s *TypeVisitor) makeSymbol(descriptor *scip.Descriptor) string {
	return symbols.FromDescriptors(s.pkg, append(s.curScope, descriptor)...)
}
