package document

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/token"
	"go/types"
	"strings"

	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

// indent is used to format struct fields.
const indent = "    "

func NewDocument(
	relative string,
	pkg *packages.Package,
	pkgSymbols *lookup.Package,
) *Document {
	return &Document{
		RelativePath: relative,
		pkg:          pkg,
		pkgSymbols:   pkgSymbols,

		docPkg: &doc.Package{},
	}
}

type Document struct {
	// Document relative path. To be used for scip.Document
	RelativePath string

	// The occurrence for `package foo` at the top of a Go file.
	//   It could be a definition or a reference, depending on the package structure.
	//   It doesn't get traversed in the same way as other parts of the tree,
	//   so we special case it here. It must get added to the occurences when
	//   creating a visitors.fileVisitor
	PackageOccurrence *scip.Occurrence

	// The package this document is contained in
	pkg *packages.Package

	// Hold information for docstrings and pretty linking
	docPkg *doc.Package

	// pkgSymbols maps positions to symbol names within
	// this document.
	pkgSymbols *lookup.Package
}

func (d *Document) GetSymbol(pos token.Pos) (string, bool) {
	return d.pkgSymbols.GetSymbol(pos)
}

// SetNewSymbol declares a new symbol and tracks it within a Document.
//
// NOTE: Does NOT emit a new occurrence
func (d *Document) SetNewSymbol(
	symbol string,
	parent ast.Node,
	ident *ast.Ident,
) {
	d.SetNewSymbolForPos(symbol, parent, ident, ident.Pos())
}

// SetNewSymbolForPos declares a new symbol and tracks it within a Document
// but allows for an override of the position. Generally speaking, you should use
// DeclareNewSymbol instead (since it will calculate the pos for most cases)
//
// NOTE: Does NOT emit a new occurrence
func (d *Document) SetNewSymbolForPos(
	symbol string,
	parent ast.Node,
	ident *ast.Ident,
	pos token.Pos,
) {
	documentation := []string{}
	if ident != nil {
		hover := d.extractHoverText(parent, ident)
		var signature, extra string
		def := d.pkg.TypesInfo.Defs[ident]
		if def != nil {
			signature, extra = typeStringForObject(def)
		}

		if signature != "" {
			documentation = append(documentation, symbols.FormatCode(signature))
		}
		if hover != "" {
			documentation = append(documentation, symbols.FormatMarkdown(hover))
		}
		if extra != "" {
			documentation = append(documentation, symbols.FormatCode(extra))
		}
	}

	d.pkgSymbols.Set(pos, &scip.SymbolInformation{
		Symbol:        symbol,
		Documentation: documentation,
		Relationships: []*scip.Relationship{},
	})
}

func (d *Document) extractHoverText(parent ast.Node, node ast.Node) string {
	switch v := node.(type) {
	case *ast.File:
		if v.Doc != nil {
			return v.Doc.Text()
		} else {
			return fmt.Sprintf("package %s", v.Name.Name)
		}
	case *ast.FuncDecl:
		return v.Doc.Text()
	case *ast.GenDecl:
		return v.Doc.Text()
	case *ast.TypeSpec:
		// Typespecs do not have the doc associated with them much
		// of the time. They are often associated with the `type`
		// token itself.
		//
		// This is why we have to pass the declaration node
		doc := v.Doc.Text()
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.ValueSpec:
		doc := v.Doc.Text()
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.Field:
		return strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
	case *ast.Ident:
		if parent != nil {
			return d.extractHoverText(nil, parent)
		}
	}

	return ""
}

// packageQualifier returns an empty string in order to remove the leading package
// name from all identifiers in the return value of types.ObjectString.
func packageQualifier(*types.Package) string { return "" }

func typeStringForType(typ types.Type) string {
	// switch ty := typ.Underlying().(type) {
	// case *types.Array:
	// 	return fmt.Sprintf("[]%s", typeStringForType(ty.Elem()))
	// }

	return typ.String()
}

func typeStringForObject(obj types.Object) (signature string, extra string) {
	switch v := obj.(type) {
	case *types.PkgName:
		return fmt.Sprintf("package %s", v.Name()), ""

	case *types.TypeName:
		return formatTypeSignature(v), formatTypeExtra(v)

	case *types.Var:
		if v.IsField() {
			// TODO(tjdevries) - make this be "(T).F" instead of "struct field F string"
			return fmt.Sprintf("struct %s", obj.String()), ""
		}

	case *types.Const:
		return fmt.Sprintf("%s = %s", types.ObjectString(v, packageQualifier), v.Val()), ""

		// TODO: We had this case in previous iterations
		// case *PkgDeclaration:
		// 	return fmt.Sprintf("package %s", v.name), ""
	}

	// Fall back to types.Object
	//    All other cases of this should be this type. We only had to implement PkgDeclaration because
	//    some fields are not exported in types.Object.
	return types.ObjectString(obj, packageQualifier), ""
}

// formatTypeSignature returns a brief description of the given struct or interface type.
func formatTypeSignature(obj *types.TypeName) string {
	switch obj.Type().Underlying().(type) {
	case *types.Struct:
		if obj.IsAlias() {
			switch obj.Type().(type) {
			case *types.Named:
				original := obj.Type().(*types.Named).Obj()
				var pkg string
				if obj.Pkg().Name() != original.Pkg().Name() {
					pkg = original.Pkg().Name() + "."
				}
				return fmt.Sprintf("type %s = %s%s", obj.Name(), pkg, original.Name())

			case *types.Struct:
				return fmt.Sprintf("type %s = struct", obj.Name())
			}
		}

		return fmt.Sprintf("type %s struct", obj.Name())
	case *types.Interface:
		return fmt.Sprintf("type %s interface", obj.Name())
	}

	return ""
}

// formatTypeExtra returns the beautified fields of the given struct or interface type.
//
// The output of `types.TypeString` puts fields of structs and interfaces on a single
// line separated by a semicolon. This method simply expands the fields to reside on
// different lines with the appropriate indentation.
func formatTypeExtra(obj *types.TypeName) string {
	extra := types.TypeString(obj.Type().Underlying(), packageQualifier)

	depth := 0
	buf := bytes.NewBuffer(make([]byte, 0, len(extra)))

outer:
	for i := 0; i < len(extra); i++ {
		switch extra[i] {
		case '"':
			for j := i + 1; j < len(extra); j++ {
				if extra[j] == '\\' {
					// skip over escaped characters
					j++
					continue
				}

				if extra[j] == '"' {
					// found non-escaped ending quote
					// write entire string unchanged, then skip to this
					// character adn continue the outer loop, which will
					// start the next iteration on the following character
					buf.WriteString(extra[i : j+1])
					i = j
					continue outer
				}
			}

			// note: we should never get down here otherwise
			// there is some illegal output from types.TypeString.

		case ';':
			buf.WriteString("\n")
			buf.WriteString(strings.Repeat(indent, depth))
			i++ // Skip following ' '

		case '{':
			// Special case empty fields so we don't insert
			// an unnecessary newline.
			if i < len(extra)-1 && extra[i+1] == '}' {
				buf.WriteString("{}")
				i++ // Skip following '}'
			} else {
				depth++
				buf.WriteString(" {\n")
				buf.WriteString(strings.Repeat(indent, depth))
			}

		case '}':
			depth--
			buf.WriteString("\n")
			buf.WriteString(strings.Repeat(indent, depth))
			buf.WriteString("}")

		default:
			buf.WriteByte(extra[i])
		}
	}

	return buf.String()
}
