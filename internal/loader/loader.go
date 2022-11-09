package loader

import (
	"github.com/sourcegraph/scip-go/internal/config"
	"golang.org/x/tools/go/packages"
)

var loadMode = packages.NeedDeps |
	packages.NeedFiles |
	packages.NeedImports |
	packages.NeedSyntax |
	packages.NeedTypes |
	packages.NeedTypesInfo |
	packages.NeedModule |
	packages.NeedName

func normalizeThisPackage(opts config.IndexOpts, pkgs []*packages.Package) {
	for _, pkg := range pkgs {
		if pkg.Module.Dir == opts.ModuleRoot {
			if pkg.Module.Version == "" {
				pkg.Module.Version = opts.ModuleVersion
			}

			if pkg.Module.Path == "" {
				pkg.Module.Path = opts.ModuleRoot
			}
		}
	}
}

func LoadPackages(opts config.IndexOpts, moduleRoot string) ([]*packages.Package, map[string]*packages.Package) {
	cfg := &packages.Config{
		Mode: loadMode,
		Dir:  moduleRoot,
		Logf: nil,

		// Only load tests for the current project.
		// This greatly reduces memory usage when loading dependencies
		Tests: true,
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		panic(err)
	}

	normalizeThisPackage(opts, pkgs)

	// TODO: Normalize the std library packages so that
	// we don't have do any special handling later on.
	//
	// This will make our lives a lot easier when reasoning
	// about packages (they will just all be loaded)
	pkgLookup := map[string]*packages.Package{
		"builtin": {
			Module: &packages.Module{
				Path:    "builtin/builtin",
				Version: "go1.19",
			},
		},
	}

	for _, pkg := range pkgs {
		ensureVersionForPackage(pkg)
		pkgLookup[pkg.Name] = pkg

		for name, imp := range pkg.Imports {
			ensureVersionForPackage(imp)
			pkgLookup[name] = imp
		}
	}

	return pkgs, pkgLookup
}

func ensureVersionForPackage(pkg *packages.Package) {
	if pkg.Module != nil {
		return
	}

	pkg.Module = &packages.Module{
		Path:    "github.com/golang/go",
		Version: "v1.19",
	}

	// fmt.Printf("Ensuring Version for Package: %s | %+v\n", pkg.PkgPath, pkg)
	// TODO: Just use the current stuff for version
	// if gomod.IsStandardlibPackge(pkg.PkgPath) {
	// 	pkg.Module = &packages.Module{
	// 		Path:    "github.com/golang/go",
	// 		Version: "v1.19",
	// 		// Main:      false,
	// 		// Indirect:  false,
	// 		// Dir:       "",
	// 		// GoMod:     "",
	// 		// GoVersion: "",
	// 		// Error:     &packages.ModuleError{},
	// 	}
	//
	// 	return
	// }

}
