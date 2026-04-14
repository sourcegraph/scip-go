package visitors

import (
	"crypto/sha256"
	"encoding/hex"
	"go/ast"
	"go/types"

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
	canon := types.TypeString(typ, func(p *types.Package) string {
		if p == nil {
			return ""
		}
		return p.Path()
	})
	sum := sha256.Sum256([]byte(canon))
	return "$anon_" + hex.EncodeToString(sum[:8])
}
