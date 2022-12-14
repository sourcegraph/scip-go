package newtypes

import (
	"go/types"

	"golang.org/x/tools/go/packages"
)

type PackageID string

func GetID(pkg *packages.Package) PackageID {
	return PackageID(pkg.PkgPath)
}

func GetFromTypesPackage(tPkg *types.Package) PackageID {
	return PackageID(tPkg.Path())
}

// func GetIDFromObject(obj types.Object) PackageID {
// 	pkg := obj.Pkg()
//
// 	// Handle Universe Scoped objs.
// 	if pkg == nil {
// 		switch v := obj.(type) {
// 		case *types.Func:
// 			switch typ := v.Type().(type) {
// 			case *types.Signature:
// 				recv := typ.Recv()
// 				universeObj := types.Universe.Lookup(recv.Type().String())
// 				if universeObj != nil {
// 					return "builtin"
// 				}
// 			}
// 		case *types.TypeName:
// 			universeObj := types.Universe.Lookup(v.Type().String())
// 			if universeObj != nil {
// 				return "builtin"
// 			}
// 		case *types.Builtin:
// 			return "builtin"
// 		case *types.Nil:
// 			return "builtin"
// 		case *types.Const:
// 			universeObj := types.Universe.Lookup(v.Type().String())
// 			if universeObj != nil {
// 				return "builtin"
// 			}
// 		}
//
// 		// Do not allow to fall through to returning pkg.Path()
// 		//
// 		// If this becomes a problem more in the future, we can just default to
// 		// returning "builtin" but as of now this handles all the cases that I
// 		// know of.
// 		// fmt.Printf("%T %+v (pkg: %s)\n", obj, obj, obj.Pkg())
// 		return "builtin"
// 	}
//
// 	// TODO: Check this out... :'(
// 	return PackageID(pkg.Path())
// }
