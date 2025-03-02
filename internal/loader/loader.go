package loader

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/output"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
	"golang.org/x/tools/go/packages"
)

type PackageLookup map[newtypes.PackageID]*packages.Package

var loadMode = packages.NeedDeps |
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
	// NOTE(id: special-test-handling): When Go code is compiled, tests
	// have special support. There are two forms of adding tests:
	//
	// - Having _test.go files with 'package mypkg' (white-box testing)
	// - Having _test.go files with 'package mypkg_test' (black-box testing)
	//
	// When using the white-box testing approach, the compiler ends up loading
	// two different packages, 'mypkg' and 'mypkg.test', with the latter also
	// including _test.go files with 'package mypkg'. Hence, if 'mypkg.test'
	// exists, then it will have the same PkgPath as 'mypkg', but it may have
	// more files.
	//
	// If it does have more files, then prefer it over the base package.
	// Ideally, we would check that we're actually using a '.test' package,
	// but the IDs are meant to be treated opaquely, so there is no reasonable
	// API to detect this.
	//
	// Apart from the special "_test" suffix, using different package names within
	// the same directory is an error, so we don't need to worry about accidentally
	// dropping information if there are unrelated packages at the same path.
	if gotPkg, ok := pkgLookup[newtypes.GetID(pkg)]; ok && len(pkg.Syntax) <= len(gotPkg.Syntax) {
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

	var panicResult any
	err = output.WithProgress("Loading Packages", func() error {
		defer func() {
			panicResult = recover()
		}()

		cfg := getConfig(moduleRoot, opts)
		patterns := opts.PackagePatterns
		if len(patterns) == 0 {
			patterns = append(patterns, "./...")
			log.Warn("No target patterns provided using default './...'")
		}
		pkgs, err := packages.Load(cfg, patterns...)
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
	})
	if err == nil && panicResult != nil {
		err = fmt.Errorf("during package loading: %v", panicResult)
	}

	if err != nil {
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

	if opts.IsGoPackagesDriverSet {
		// As of golang.org/x/tools v0.26, the JSON object for representing packages
		// does not have a Module field:
		//   https://github.com/golang/tools/blob/v0.26.0/go/packages/packages.go#L609-L627
		// which means that cross-repo navigation may not work out-of-the-box
		// for 3rd party build systems like Bazel, Buck2, Please etc. using GOPACKAGESDRIVER
		pkg.Module = &packages.Module{
			Path:    ".",
			Version: ".",
		}

		if opts.ModulePath != "" {
			pkg.Module.Path = opts.ModulePath
		}

		if opts.ModuleVersion != "" {
			pkg.Module.Version = opts.ModuleVersion
		}

		return pkg
	}

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
			log.Error(handler.ErrOrPanic(
				"Unknown version for userland package: %s %s",
				pkg.Module.Path,
				opts.ModulePath,
			))

			pkg.Module.Version = "."
		}
	} else if module.IsPseudoVersion(pkg.Module.Version) {
		// Unpublished versions of dependencies have pseudo-versions in go.mod.
		// When the dependency itself is indexed, only the revision will be used.
		// For correct cross-repo navigation to such dependencies, only use
		// the revision from a pseudo-version.
		rev, err := module.PseudoVersionRev(pkg.Module.Version)
		if err != nil {
			// Only panic when running in debug mode.
			log.Error(handler.ErrOrPanic(
				"Unable to find rev from pseudo-version: %s %s",
				pkg.Module.Path,
				pkg.Module.Version,
			))
		} else {
			pkg.Module.Version = rev
		}
	} else if build := semver.Build(pkg.Module.Version); build != "" {
		// The revision can also have build metadata following a `+`. Drop that,
		// similar to official Go tooling. (https://go.dev/ref/mod#versions)
		// > The build metadata suffix is ignored for the purpose of comparing versions
		pkg.Module.Version = strings.TrimSuffix(pkg.Module.Version, build)
	}

	return pkg
}
