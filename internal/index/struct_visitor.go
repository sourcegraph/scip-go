package index

import (
	"fmt"
	"go/ast"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitFieldsInFile(doc *document.Document, pkg *packages.Package, file *ast.File) {
	visitor := FieldVisitor{
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

// FieldVisitor collects the all the information for top-level structs
// that can be imported by any other file (they do not have to be exported).
//
// For example, a struct `myStruct` can be imported by other files in the same
// packages. So we need to make those field names global (we only have global
// or file-local).
type FieldVisitor struct {
	doc *document.Document
	pkg *packages.Package

	curScope []*scip.Descriptor
	curDecl  *ast.GenDecl
}

func (v FieldVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.GenDecl:
		// Current declaration is required for some documentation parsing.
		// So we have to keep this here with us as we traverse more deeply
		v.curDecl = node
		return v

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

		v.doc.DeclareNewSymbol(
			symbols.FromDescriptors(v.pkg, v.curScope...),
			v.curDecl,
			node.Name,
		)

		ast.Walk(v, node.Type)
	case *ast.Field:
		// I think the only case of this is embedded fields.
		if len(node.Names) == 0 {
			name := v.getIdentOfTypeExpr(node.Type)
			embeddedSymbol := v.makeSymbol(&scip.Descriptor{
				Name:   name.Name,
				Suffix: scip.Descriptor_Term,
			})

			// In this odd scenario, the definition is at embedded field level,
			// not wherever the name is. So that messes up our lookup table.
			v.doc.DeclareNewSymbolForPos(embeddedSymbol, node, name, node.Pos())
		} else {
			for _, name := range node.Names {
				v.doc.DeclareNewSymbol(v.makeSymbol(&scip.Descriptor{
					Name:   name.Name,
					Suffix: scip.Descriptor_Term,
				}), nil, name)

				switch node.Type.(type) {
				case *ast.StructType, *ast.InterfaceType:
					// Current scope is now embedded in the anonymous struct
					//   So we walk the rest of the type expression and save
					//   the nested names
					v.curScope = append(v.curScope, &scip.Descriptor{
						Name:   name.Name,
						Suffix: scip.Descriptor_Term,
					})

					ast.Walk(v, node.Type)

					v.curScope = v.curScope[:len(v.curScope)-1]
				}
			}
		}
	}

	return nil
}

// Implements ast.Visitor
var _ ast.Visitor = &FieldVisitor{}

func (s *FieldVisitor) getIdentOfTypeExpr(ty ast.Expr) *ast.Ident {
	switch ty := ty.(type) {
	case *ast.Ident:
		return ty
	case *ast.SelectorExpr:
		return ty.Sel
	case *ast.StarExpr:
		return s.getIdentOfTypeExpr(ty.X)
	default:
		panic(fmt.Sprintf("Unhandled unamed struct field: %T %+v", ty, ty))
	}
}

func (s *FieldVisitor) makeSymbol(descriptor *scip.Descriptor) string {
	return symbols.FromDescriptors(s.pkg, append(s.curScope, descriptor)...)
}
