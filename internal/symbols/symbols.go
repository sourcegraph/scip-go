package symbols

import (
	"fmt"
	"go/token"
	"go/types"
	"sync"

	"github.com/scip-code/scip/bindings/go/scip"
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

var synthesisCache sync.Map // map[types.Object]synthesisResult

type synthesisResult struct {
	sym string
	ok  bool
}

// SynthesizeFromObject constructs a SCIP symbol string for a types.Object
// that belongs to a dependency package (no AST available). It replicates the
// same descriptor chain that the AST-based visitors would produce.
func SynthesizeFromObject(pkg *packages.Package, obj types.Object) (string, bool) {
	if v, ok := synthesisCache.Load(obj); ok {
		r := v.(synthesisResult)
		return r.sym, r.ok
	}
	sym, ok := synthesizeFromObject(pkg, obj)
	synthesisCache.Store(obj, synthesisResult{sym, ok})
	return sym, ok
}

func synthesizeFromObject(pkg *packages.Package, obj types.Object) (string, bool) {
	if pkg.Module == nil {
		return "", false
	}

	pkgPath := obj.Pkg().Path()
	descriptors := []*scip.Descriptor{
		{Name: pkgPath, Suffix: scip.Descriptor_Namespace},
	}

	switch obj := obj.(type) {
	case *types.Func:
		sig, _ := obj.Type().(*types.Signature)
		if sig != nil && sig.Recv() != nil {
			recvName := receiverTypeName(sig.Recv().Type())
			if recvName == "" {
				return "", false
			}
			descriptors = append(descriptors,
				&scip.Descriptor{Name: recvName, Suffix: scip.Descriptor_Type},
				&scip.Descriptor{Name: obj.Name(), Suffix: scip.Descriptor_Method},
			)
		} else {
			descriptors = append(descriptors,
				&scip.Descriptor{Name: obj.Name(), Suffix: scip.Descriptor_Method},
			)
		}

	case *types.TypeName:
		descriptors = append(descriptors,
			&scip.Descriptor{Name: obj.Name(), Suffix: scip.Descriptor_Type},
		)

	case *types.Var:
		if obj.IsField() {
			typeName := findFieldOwner(obj)
			if typeName == "" {
				return "", false
			}
			descriptors = append(descriptors,
				&scip.Descriptor{Name: typeName, Suffix: scip.Descriptor_Type},
				&scip.Descriptor{Name: obj.Name(), Suffix: scip.Descriptor_Term},
			)
		} else {
			descriptors = append(descriptors,
				&scip.Descriptor{Name: obj.Name(), Suffix: scip.Descriptor_Term},
			)
		}

	case *types.Const:
		descriptors = append(descriptors,
			&scip.Descriptor{Name: obj.Name(), Suffix: scip.Descriptor_Term},
		)

	default:
		return "", false
	}

	sym := scip.VerboseSymbolFormatter.FormatSymbol(&scip.Symbol{
		Scheme: "scip-go",
		Package: &scip.Package{
			Manager: "gomod",
			Name:    pkg.Module.Path,
			Version: pkg.Module.Version,
		},
		Descriptors: descriptors,
	})
	return sym, true
}

// receiverTypeName extracts the base named type from a receiver type,
// peeling off pointers and aliases.
func receiverTypeName(t types.Type) string {
	for {
		switch v := types.Unalias(t).(type) {
		case *types.Pointer:
			t = v.Elem()
		case *types.Named:
			return v.Obj().Name()
		default:
			return ""
		}
	}
}

var fieldOwnerCache sync.Map // map[*types.Var]string

// buildFieldIndex builds a field→owner-name lookup for all struct types
// in a package scope, caching the result per package.
var fieldIndexCache sync.Map // map[*types.Package]map[*types.Var]string

func getFieldIndex(pkg *types.Package) map[*types.Var]string {
	if v, ok := fieldIndexCache.Load(pkg); ok {
		return v.(map[*types.Var]string)
	}
	idx := map[*types.Var]string{}
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		tn, ok := obj.(*types.TypeName)
		if !ok {
			continue
		}
		named, ok := tn.Type().(*types.Named)
		if !ok {
			continue
		}
		st, ok := named.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		for i := range st.NumFields() {
			idx[st.Field(i)] = tn.Name()
		}
	}
	fieldIndexCache.Store(pkg, idx)
	return idx
}

// findFieldOwner finds the name of the named type that directly declares a
// struct field.
func findFieldOwner(field *types.Var) string {
	if v, ok := fieldOwnerCache.Load(field); ok {
		return v.(string)
	}
	pkg := field.Pkg()
	if pkg == nil {
		return ""
	}
	idx := getFieldIndex(pkg)
	result := idx[field]
	fieldOwnerCache.Store(field, result)
	return result
}

func FormatCode(v string) string {
	if v == "" {
		return ""
	}

	return fmt.Sprintf("```go\n%s\n```", v)
}
