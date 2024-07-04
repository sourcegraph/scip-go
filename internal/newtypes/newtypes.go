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
