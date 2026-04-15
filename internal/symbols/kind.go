package symbols

import (
	"go/types"

	"github.com/scip-code/scip/bindings/go/scip"
)

// KindForObject maps a Go types.Object to the corresponding SCIP SymbolInformation_Kind.
func KindForObject(obj types.Object) scip.SymbolInformation_Kind {
	if obj == nil {
		return scip.SymbolInformation_UnspecifiedKind
	}

	switch obj := obj.(type) {
	case *types.Func:
		sig, ok := obj.Type().(*types.Signature)
		if !ok {
			return scip.SymbolInformation_Function
		}
		if sig.Recv() == nil {
			return scip.SymbolInformation_Function
		}
		if types.IsInterface(sig.Recv().Type()) {
			return scip.SymbolInformation_MethodSpecification
		}
		return scip.SymbolInformation_Method

	case *types.TypeName:
		if obj.IsAlias() {
			return scip.SymbolInformation_TypeAlias
		}
		switch obj.Type().Underlying().(type) {
		case *types.Struct:
			return scip.SymbolInformation_Struct
		case *types.Interface:
			return scip.SymbolInformation_Interface
		case *types.TypeParam:
			return scip.SymbolInformation_TypeParameter
		default:
			return scip.SymbolInformation_Type
		}

	case *types.Var:
		if obj.IsField() {
			return scip.SymbolInformation_Field
		}
		return scip.SymbolInformation_Variable

	case *types.Const:
		return scip.SymbolInformation_Constant

	case *types.PkgName:
		return scip.SymbolInformation_Package

	default:
		return scip.SymbolInformation_UnspecifiedKind
	}
}
