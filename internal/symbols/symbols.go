package symbols

import (
	"fmt"
	"go/token"
	"go/types"

	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func FromDescriptors(pkg *packages.Package, descriptors ...*scip.Descriptor) string {
	if pkg.Module == nil {
		panic(fmt.Sprintf("Failed to find package for: %+v\n", pkg.PkgPath))
	}

	return scip.VerboseSymbolFormatter.FormatSymbol(&scip.Symbol{
		Scheme: "scip-go",
		Package: &scip.Package{
			Manager: "gomod",
			Name:    pkg.Module.Path,
			Version: pkg.Module.Version,
		},
		Descriptors: descriptors,
	})
}

func RangeFromName(position token.Position, name string, adjust bool) []int32 {
	var adjustment int32 = 0
	if adjust {
		adjustment = 1
	}

	line := int32(position.Line - 1)
	column := int32(position.Column - 1)
	n := int32(len(name))

	return []int32{line, column + adjustment, column + n + adjustment}
}

func FormatCode(v string) string {
	if v == "" {
		return ""
	}

	return fmt.Sprintf("```go\n%s\n```", v)
}

// GetSymbolKind 根据 Go 类型对象返回对应的 SCIP 符号类型
func GetSymbolKind(obj types.Object) scip.SymbolInformation_Kind {
	switch obj := obj.(type) {
	case *types.PkgName:
		return scip.SymbolInformation_Namespace
	case *types.Const:
		return scip.SymbolInformation_Constant
	case *types.Var:
		if obj.IsField() {
			return scip.SymbolInformation_Field
		}
		return scip.SymbolInformation_Variable
	case *types.Func:
		sig := obj.Type().(*types.Signature)
		if sig.Recv() != nil {
			return scip.SymbolInformation_Method
		}
		return scip.SymbolInformation_Function
	case *types.TypeName:
		switch obj.Type().Underlying().(type) {
		case *types.Struct:
			return scip.SymbolInformation_Class
		case *types.Interface:
			return scip.SymbolInformation_Interface
		default:
			return scip.SymbolInformation_Type
		}
	default:
		return scip.SymbolInformation_UnspecifiedKind
	}
}
