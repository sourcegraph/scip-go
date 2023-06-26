package loader

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/output"
	"golang.org/x/mod/modfile"
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

func getConfig(root string, opts config.IndexOpts) *packages.Config {
	// TODO: Hacks to get the config out...
	Config = &packages.Config{
		Mode: loadMode,
		Dir:  root,
		Logf: output.Logf,

		// Only load tests for the current project.
		// This greatly reduces memory usage when loading dependencies
		Tests: !opts.SkipTests,
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

func LoadPackages(
	opts config.IndexOpts,
	moduleRoot string,
) (pkgLookup PackageLookup, projectPackages PackageLookup, err error) {
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
		cfg := getConfig(moduleRoot, opts)
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
	//  PkgPath = net/http
	//  -> net
	//  -> true
	//
	//  PkgPath = github.com/sourcegraph/scip-go/...
	//  -> github.com/
	//  -> false
	base := strings.Split(pkg.PkgPath, "/")[0]
	if _, ok := stdPackages[base]; ok {
		return ok
	}

	noTestPackage := strings.Replace(base, "_test", "", -1)
	if _, ok := stdPackages[noTestPackage]; ok {
		return ok
	}

	noTestPsuedoPackage := strings.Replace(base, ".test", "", -1)
	if _, ok := stdPackages[noTestPsuedoPackage]; ok {
		return ok
	}

	return false
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

	// Follow replaced modules
	if pkg.Module.Replace != nil {
		pkg.Module = pkg.Module.Replace

		// Local replaces for local files can have this happen,
		// so short circuit the check here (the following versions should not be able to fail)
		if pkg.Module.Version == "" {
			pkg.Module.Version = opts.ModuleVersion
		}
	}

	// Replace **local** directives with the resolved go package.
	// Attempt to parse the go.mod file (with the builtin `modfile` package) and
	// then update the module path appropriately
	if strings.HasPrefix(pkg.Module.Path, ".") {
		if pkg.Module.GoMod != "" {
			contents, err := ioutil.ReadFile(pkg.Module.GoMod)

			if err != nil {
				handler.ErrOrPanic("Failed to read go mod file: %s", err)
			} else {
				parsed, err := modfile.ParseLax(pkg.Module.GoMod, contents, nil)
				if err != nil {
					handler.ErrOrPanic("Failed to parse go mod file: %s", err)
				}

				output.Logf("[scip.loader] Replacing module path: '%s' with '%s'", pkg.Module.Path, parsed.Module.Mod.Path)
				pkg.Module.Path = parsed.Module.Mod.Path

				// If we have a version specified in this go.mod, we'll use that.
				// Otherwise we'll fall back to whatever the version was previous set to.
				if parsed.Module.Mod.Version != "" {
					output.Logf("[scip.loader] Replacing module version: '%s' with '%s'", pkg.Module.Version, parsed.Module.Mod.Version)
					pkg.Module.Version = parsed.Module.Mod.Version
				}
			}
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
