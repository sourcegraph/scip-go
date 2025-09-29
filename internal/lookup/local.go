package lookup

import (
	"go/types"
	"strings"
)

// Local contains information about a local symbol
type Local struct {
	Symbol string
	Obj    types.Object
}

// SignatureText builds a concise signature string for this local symbol.
func (l *Local) SignatureText() string {
	if l.Obj == nil {
		return ""
	}

	var parts []string

	// local symbol can only represent const or var
	switch l.Obj.(type) {
	case *types.Const:
		parts = append(parts, "const")
	case *types.Var:
		parts = append(parts, "var")
	}

	if name := l.Obj.Name(); name != "" {
		parts = append(parts, name)
	}

	if t := l.Obj.Type(); t != nil {
		if ts := t.String(); ts != "" {
			parts = append(parts, ts)
		}
	}

	return strings.Join(parts, " ")
}
