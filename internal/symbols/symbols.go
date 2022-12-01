package symbols

import (
	"fmt"
	"go/types"

	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

// FromObject generates a type from a struct.
//
// At the moment, I only want this to happen from builtin objects (everything
// else should have been traversed in the first pass of the indexer).
//
// The main reason for that is this would not properly handle things like
// nested struct fields, etc.
func FromObject(pkg *packages.Package, obj types.Object) (string, bool) {
	if pkg == nil {
		panic("Invalid, must pass a package")
	}

	if pkg.Name != "builtin" {
		panic(fmt.Sprintf("Can only generate from object if pkg is builtin %s", pkg.Name))
	}

	desc := []*scip.Descriptor{
		{Name: makeMonikerPackage(obj), Suffix: scip.Descriptor_Package},
	}
	return FromDescriptors(pkg, append(desc, scipDescriptors(obj)...)...), true
}

func makeMonikerPackage(obj types.Object) string {
	var pkgName string
	if v, ok := obj.(*types.PkgName); ok {
		// gets the full path of the package name, rather than just the name.
		// So instead of "http", it will return "net/http"
		pkgName = v.Imported().Path()
	} else {
		pkgName = PkgPathFromObject(obj)
	}

	// return gomod.NormalizeMonikerPackage(pkgName)
	// TODO normalize name
	return pkgName
}

func scipDescriptors(obj types.Object) []*scip.Descriptor {
	switch obj := obj.(type) {
	case *types.Func:
		return []*scip.Descriptor{
			{Name: obj.Name(), Suffix: scip.Descriptor_Method},
		}
	case *types.Var:
		if obj.IsField() {
			panic("Don't do fields")
		}

		return []*scip.Descriptor{
			{Name: obj.Name(), Suffix: scip.Descriptor_Term},
		}
	// case *types.Const:
	// 	return []*scip.Descriptor{
	// 		{Name: obj.Name(), Suffix: scip.Descriptor_Term},
	// 	}
	case *types.TypeName:
		return []*scip.Descriptor{
			{Name: obj.Name(), Suffix: scip.Descriptor_Type},
		}
	// case *types.PkgName:
	// 	return []*scip.Descriptor{
	// 		{Name: obj.Name(), Suffix: scip.Descriptor_Namespace},
	// 	}
	case *types.Builtin:
		return []*scip.Descriptor{
			{Name: obj.Name(), Suffix: scip.Descriptor_Term},
		}

	default:
		panic(fmt.Sprintf("unknown scip descriptor for type: |%T| %+v\n", obj, obj))
	}

	// return []*scip.Descriptor{}
}

func FromDescriptors(pkg *packages.Package, descriptors ...*scip.Descriptor) string {
	if pkg.Module == nil {
		panic(fmt.Sprintf("Failed to find package for: %+v\n", pkg.PkgPath))
	}

	return scip.VerboseSymbolFormatter.FormatSymbol(&scip.Symbol{
		Scheme: "scip-go",
		Package: &scip.Package{
			Manager: "gomod",
			// TODO: We might not have a dep, so we should handle that
			Name:    pkg.Module.Path,
			Version: pkg.Module.Version,
		},
		Descriptors: descriptors,
	})
}

func PkgPathFromObject(obj types.Object) string {
	pkg := obj.Pkg()

	// Handle Universe Scoped objs.
	if pkg == nil {
		switch v := obj.(type) {
		case *types.Func:
			switch typ := v.Type().(type) {
			case *types.Signature:
				recv := typ.Recv()
				universeObj := types.Universe.Lookup(recv.Type().String())
				if universeObj != nil {
					return "builtin"
				}
			}
		case *types.TypeName:
			universeObj := types.Universe.Lookup(v.Type().String())
			if universeObj != nil {
				return "builtin"
			}
		case *types.Builtin:
			return "builtin"
		case *types.Nil:
			return "builtin"
		case *types.Const:
			universeObj := types.Universe.Lookup(v.Type().String())
			if universeObj != nil {
				return "builtin"
			}
		}

		// Do not allow to fall through to returning pkg.Path()
		//
		// If this becomes a problem more in the future, we can just default to
		// returning "builtin" but as of now this handles all the cases that I
		// know of.
		// fmt.Printf("%T %+v (pkg: %s)\n", obj, obj, obj.Pkg())
		return "builtin"
		// panic("Unhandled nil obj.Pkg()")
		// return "builtin"
	}

	return pkg.Path()
}
