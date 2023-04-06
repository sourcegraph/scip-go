package loader

import (
	"fmt"
	"strings"

	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/output"
	"golang.org/x/tools/go/packages"
)

type PackageLookup map[newtypes.PackageID]*packages.Package

var loadMode = packages.NeedDeps |
	packages.NeedFiles |
	packages.NeedImports |
	packages.NeedSyntax |
	packages.NeedTypes |
	packages.NeedTypesInfo |
	packages.NeedModule |
	packages.NeedName

var Config = &packages.Config{}

func makeConfig(root string) *packages.Config {
	// TODO: Hacks to get the config out...
	Config = &packages.Config{
		Mode: loadMode,
		Dir:  root,
		Logf: nil,

		// Only load tests for the current project.
		// This greatly reduces memory usage when loading dependencies
		Tests: true,
	}

	return Config
}

func addImportsToPkgs(pkgLookup PackageLookup, opts *config.IndexOpts, pkg *packages.Package) {
	if _, ok := pkgLookup[newtypes.GetID(pkg)]; ok {
		return
	}

	normalizePackage(opts, pkg)
	pkgLookup[newtypes.GetID(pkg)] = pkg

	for _, imp := range pkg.Imports {
		addImportsToPkgs(pkgLookup, opts, imp)
	}
}

func LoadPackages(opts config.IndexOpts, moduleRoot string) (pkgLookup PackageLookup, projectPackages PackageLookup, err error) {
	// Force a module version, even if it's just a dot for non-cross repo look ups.
	if opts.ModuleVersion == "" {
		opts.ModuleVersion = "."
	}

	pkgLookup = make(PackageLookup)
	pkgLookup["builtin"] = &packages.Package{
		Name:    "builtin",
		PkgPath: "builtin",
		Module: &packages.Module{
			Path:    "github.com/golang/go/src/builtin",
			Version: opts.GoStdlibVersion,
		},
	}

	projectPackages = make(PackageLookup)

	if err := output.WithProgress("Loading Packages", func() error {
		cfg := makeConfig(moduleRoot)
		pkgs, err := packages.Load(cfg, "./...")
		if err != nil {
			return err
		}

		for _, pkg := range pkgs {
			addImportsToPkgs(pkgLookup, &opts, pkg)
		}

		for _, pkg := range pkgs {
			projectPackages[newtypes.GetID(pkg)] = pkg
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

	return projectPackages, pkgLookup, nil
}

func IsStandardLib(pkg *packages.Package) bool {
	// for example:
	//	PkgPath = net/http
	//	-> net
	//	-> true
	//
	//	PkgPath = github.com/sourcegraph/scip-go/...
	//	-> github.com/
	//	-> false
	base := strings.Split(pkg.PkgPath, "/")[0]
	_, ok := stdPackages[base]
	return ok
}

func normalizePackage(opts *config.IndexOpts, pkg *packages.Package) *packages.Package {
	// Name string = "pentimento"
	// PkgPath string = "github.com/efritz/pentimento"
	// Module:
	//		Path string = "github.com/efritz/pentimento"
	//		Version string = "v0.0.0-20190429011147-ade47d831101"

	if IsStandardLib(pkg) || opts.IsIndexingStdlib {
		pkg.Module = &packages.Module{
			Path:    "github.com/golang/go/src",
			Version: opts.GoStdlibVersion,
		}

		// When indexing the standard library, all packages are prefixed with `std/`.
		//
		// We strip that to standardize all the libraries to make sure we are able to jump-to-definition
		// of the standard library.
		pkg.PkgPath = strings.TrimPrefix(pkg.PkgPath, "std/")
	} else {
		if pkg.Module == nil {
			panic(fmt.Sprintf(
				"Should not be possible to have nil module for userland package: %s %s",
				pkg,
				pkg.PkgPath,
			))
		}

	}

	// TODO: Handle `./lib` style
	// TODO: Ensure that we copy version correclty

	// Follow replaced modules
	if pkg.Module.Replace != nil {
		pkg.Module = pkg.Module.Replace

		// Local replaces for local files can have this happen,
		// so short circuit the check here (the following versions should not be able to fail)
		if pkg.Module.Version == "" {
			pkg.Module.Version = opts.ModuleVersion
		}
	}

	if pkg.Module.Path == "" {
		pkg.Module.Path = "."
	}

	if pkg.Module.Version == "" {
		if pkg.Module.Path == opts.ModulePath {
			pkg.Module.Version = opts.ModuleVersion
		} else {
			// Only panic when running in debug mode.
			fmt.Println(handler.ErrOrPanic(
				"Unknown version for userland package: %s %s",
				pkg.Module.Path,
				opts.ModulePath,
			))

			pkg.Module.Version = "."
		}
	}

	return pkg
}
