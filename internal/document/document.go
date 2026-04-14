package document

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/format"
	"go/token"
	"go/types"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/sourcegraph/scip-go/internal/lookup"
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

// SetSymbolInformation registers a pre-built SymbolInformation at the given position.
func (d *Document) SetSymbolInformation(
	pos token.Pos, info *scip.SymbolInformation,
) {
	d.pkgSymbols.Set(pos, info)
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
	var documentation []string
	var sigDoc *scip.Document
	if ident != nil {
		if def := d.pkg.TypesInfo.Defs[ident]; def != nil {
			if signature := typeStringForObject(def); signature != "" {
				sigDoc = &scip.Document{
					Language: "go",
					Text:     signature,
				}
			}
		}
		if hover := d.extractHoverText(parent, ident); hover != "" {
			documentation = append(documentation, hover)
		}
	}

	d.pkgSymbols.Set(pos, &scip.SymbolInformation{
		Symbol:                 symbol,
		Documentation:          documentation,
		SignatureDocumentation: sigDoc,
		Relationships:          []*scip.Relationship{},
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
		doc := strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.Field:
		return strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
	case *ast.Ident:
		if genDecl, ok := parent.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if s.Name == v {
						return d.extractHoverText(parent, s)
					}
				case *ast.ValueSpec:
					for _, name := range s.Names {
						if name == v {
							return d.extractHoverText(parent, s)
						}
					}
				}
			}
		}
		if parent != nil {
			return d.extractHoverText(nil, parent)
		}
	}

	return ""
}

// packageQualifier returns an empty string in order to remove the leading package
// name from all identifiers in the return value of types.ObjectString.
func packageQualifier(*types.Package) string { return "" }

func typeStringForObject(obj types.Object) string {
	switch v := obj.(type) {
	case *types.PkgName:
		return fmt.Sprintf("package %s", v.Name())

	case *types.TypeName:
		return formatTypeDeclaration(v)

	case *types.Var:
		if v.IsField() {
			// TODO(tjdevries) - make this be "(T).F" instead of "struct field F string"
			return fmt.Sprintf("struct %s", quotedTagsToBacktick(obj.String()))
		}

	case *types.Const:
		return fmt.Sprintf("%s = %s", types.ObjectString(v, packageQualifier), v.Val())
	}

	return types.ObjectString(obj, packageQualifier)
}

var loggedGODEBUGWarning sync.Once

// formatTypeDeclaration returns the full type declaration string for a type,
// e.g. "type T struct{}", "type I interface { ... }", "type U = T", "type Z int32".
func formatTypeDeclaration(obj *types.TypeName) string {
	if obj.IsAlias() {
		return formatAliasDeclaration(obj)
	}

	return fmt.Sprintf("type %s %s", obj.Name(), expandTypeExpr(obj.Type().Underlying()))
}

// qualifiedName returns the name of original, prefixed with its package name
// if it differs from obj's package. Handles nil packages (builtins).
func qualifiedName(obj, original types.Object) string {
	objPkg := obj.Pkg()
	origPkg := original.Pkg()
	if objPkg == nil || origPkg == nil || objPkg.Name() == origPkg.Name() {
		return original.Name()
	}
	return origPkg.Name() + "." + original.Name()
}

// formatAliasDeclaration returns the type declaration for alias types.
func formatAliasDeclaration(obj *types.TypeName) string {
	switch ty := obj.Type().(type) {
	case *types.Alias:
		switch rhs := ty.Rhs().(type) {
		case *types.Alias:
			return fmt.Sprintf("type %s = %s", obj.Name(), qualifiedName(obj, rhs.Obj()))
		case *types.Named:
			return fmt.Sprintf("type %s = %s", obj.Name(), qualifiedName(obj, rhs.Obj()))
		default:
			return fmt.Sprintf("type %s = %s", obj.Name(), expandTypeExpr(rhs))
		}
	default:
		if val := os.Getenv("GODEBUG"); strings.Contains(val, "gotypealias=0") {
			loggedGODEBUGWarning.Do(func() {
				slog.Warn(
					"Running with GODEBUG=gotypealias=0, this may cause incorrect hover docs")
			})
		} else {
			slog.Warn(
				"IsAlias() is true but Type() is not Alias; please report this as a bug",
				"obj", obj.String(), "obj.Type()", ty.String())
		}
	}

	// Fallback for when GODEBUG=gotypealias=0 or unexpected types.
	return fmt.Sprintf("type %s %s", obj.Name(), expandTypeExpr(obj.Type().Underlying()))
}

// expandTypeExpr renders a type expression, formatting struct and interface
// types with aligned fields using go/format.
func expandTypeExpr(t types.Type) string {
	raw := types.TypeString(t, packageQualifier)

	switch t.(type) {
	case *types.Struct, *types.Interface:
		if formatted, err := formatGoDecl("type _ " + raw); err == nil {
			return strings.TrimPrefix(formatted, "type _ ")
		}
	}

	return raw
}

// formatGoDecl formats a Go type declaration using go/format, returning
// the formatted result with tabs replaced by spaces.
func formatGoDecl(decl string) (string, error) {
	src := "package p\n\n" + decl + "\n"
	formatted, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}
	result := strings.TrimPrefix(string(formatted), "package p\n\n")
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "\t", indent)
	return result, nil
}

// quotedTagsToBacktick replaces double-quoted struct tag strings with
// backtick-quoted equivalents in the output of types.ObjectString.
func quotedTagsToBacktick(s string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	for i := 0; i < len(s); i++ {
		if s[i] != '"' {
			buf.WriteByte(s[i])
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[j] == '\\' {
				j++
				continue
			}
			if s[j] == '"' {
				quoted := s[i : j+1]
				if unquoted, err := strconv.Unquote(quoted); err == nil {
					buf.WriteByte('`')
					buf.WriteString(unquoted)
					buf.WriteByte('`')
				} else {
					buf.WriteString(quoted)
				}
				i = j
				break
			}
		}
	}
	return buf.String()
}
