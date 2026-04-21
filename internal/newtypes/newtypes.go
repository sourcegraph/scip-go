package newtypes

import (
	"go/types"

	"golang.org/x/tools/go/packages"
)

// PackageID is a Go package's import path (e.g. "fmt", "encoding/json",
// "github.com/scip-code/scip-go/internal/loader").
type PackageID string

func GetID(pkg *packages.Package) PackageID {
	return PackageID(pkg.PkgPath)
}

func GetFromTypesPackage(tPkg *types.Package) PackageID {
	return PackageID(tPkg.Path())
}
