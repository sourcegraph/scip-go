package visitors

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

// anonTypeKey computes a canonical, stable key for an anonymous struct or
// interface type. Two anonymous types that are identical (per Go's type
// identity rules) will produce the same key.
//
// The key is derived from go/types (not the AST) so it correctly handles
// aliases, generics, field order, struct tags, and embedded fields.
func anonTypeKey(pkg *packages.Package, expr ast.Expr) string {
	typ := pkg.TypesInfo.TypeOf(expr)
	if typ == nil {
		return ""
	}
	canon := canonicalType(typ, nil)
	sum := sha256.Sum256([]byte(canon))
	return "$anon_" + hex.EncodeToString(sum[:8])
}

// canonicalType produces a stable string representation of a type suitable
// for hashing. It preserves field order, struct tags, embeddedness, and
// distinguishes named types from their underlying types.
func canonicalType(t types.Type, seen map[*types.Named]bool) string {
	if seen == nil {
		seen = map[*types.Named]bool{}
	}

	switch t := t.(type) {
	case *types.Struct:
		return canonicalStruct(t, seen)
	case *types.Interface:
		return canonicalInterface(t, seen)
	case *types.Named:
		pkg := t.Obj().Pkg()
		pkgPath := ""
		if pkg != nil {
			pkgPath = pkg.Path()
		}
		if seen[t] {
			return fmt.Sprintf("@%s.%s", pkgPath, t.Obj().Name())
		}
		seen[t] = true
		var b strings.Builder
		b.WriteString(pkgPath)
		b.WriteByte('.')
		b.WriteString(t.Obj().Name())
		if targs := t.TypeArgs(); targs != nil && targs.Len() > 0 {
			b.WriteByte('[')
			for i := 0; i < targs.Len(); i++ {
				if i > 0 {
					b.WriteByte(',')
				}
				b.WriteString(canonicalType(targs.At(i), seen))
			}
			b.WriteByte(']')
		}
		return b.String()
	case *types.Alias:
		return canonicalType(types.Unalias(t), seen)
	case *types.Pointer:
		return "*" + canonicalType(t.Elem(), seen)
	case *types.Slice:
		return "[]" + canonicalType(t.Elem(), seen)
	case *types.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), canonicalType(t.Elem(), seen))
	case *types.Map:
		return fmt.Sprintf("map[%s]%s", canonicalType(t.Key(), seen), canonicalType(t.Elem(), seen))
	case *types.Chan:
		var dir string
		switch t.Dir() {
		case types.SendRecv:
			dir = "chan "
		case types.SendOnly:
			dir = "chan<- "
		case types.RecvOnly:
			dir = "<-chan "
		}
		return dir + canonicalType(t.Elem(), seen)
	case *types.Signature:
		return canonicalSignature(t, seen)
	case *types.Basic:
		return t.Name()
	case *types.TypeParam:
		return fmt.Sprintf("$%d", t.Index())
	default:
		return t.String()
	}
}

func canonicalStruct(t *types.Struct, seen map[*types.Named]bool) string {
	var b strings.Builder
	b.WriteString("struct{")
	for i := 0; i < t.NumFields(); i++ {
		if i > 0 {
			b.WriteByte(';')
		}
		f := t.Field(i)
		if f.Embedded() {
			b.WriteString("~")
		}
		if !f.Exported() && f.Pkg() != nil {
			b.WriteString(f.Pkg().Path())
			b.WriteByte('.')
		}
		b.WriteString(f.Name())
		b.WriteByte(' ')
		b.WriteString(canonicalType(f.Type(), seen))
		if tag := t.Tag(i); tag != "" {
			b.WriteString(" `")
			b.WriteString(tag)
			b.WriteByte('`')
		}
	}
	b.WriteByte('}')
	return b.String()
}

func canonicalInterface(t *types.Interface, seen map[*types.Named]bool) string {
	var b strings.Builder
	b.WriteString("interface{")
	for i := 0; i < t.NumMethods(); i++ {
		if i > 0 {
			b.WriteByte(';')
		}
		m := t.Method(i)
		b.WriteString(m.Name())
		b.WriteString(canonicalSignature(m.Type().(*types.Signature), seen))
	}
	for i := 0; i < t.NumEmbeddeds(); i++ {
		if i > 0 || t.NumMethods() > 0 {
			b.WriteByte(';')
		}
		b.WriteString(canonicalType(t.EmbeddedType(i), seen))
	}
	b.WriteByte('}')
	return b.String()
}

func canonicalSignature(t *types.Signature, seen map[*types.Named]bool) string {
	var b strings.Builder
	b.WriteString("func(")
	params := t.Params()
	for i := 0; i < params.Len(); i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(canonicalType(params.At(i).Type(), seen))
	}
	b.WriteByte(')')
	results := t.Results()
	if results.Len() > 0 {
		b.WriteByte('(')
		for i := 0; i < results.Len(); i++ {
			if i > 0 {
				b.WriteByte(',')
			}
			b.WriteString(canonicalType(results.At(i).Type(), seen))
		}
		b.WriteByte(')')
	}
	return b.String()
}
