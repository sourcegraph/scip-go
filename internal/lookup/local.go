package lookup

import "go/types"

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
	if pkgName, ok := l.Obj.(*types.PkgName); ok {
		if imported := pkgName.Imported(); imported != nil {
			return "import " + pkgName.Name() + " " + imported.Path()
		}
		return "import " + pkgName.Name()
	}
	return types.ObjectString(l.Obj, func(*types.Package) string { return "" })
}
