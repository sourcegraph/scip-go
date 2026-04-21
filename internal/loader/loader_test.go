package loader

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/scip-code/scip-go/internal/config"
	"golang.org/x/mod/module"
	"golang.org/x/tools/go/packages"
)

func TestBuiltinFormat(t *testing.T) {
	wd, _ := os.Getwd()
	root, _ := filepath.Abs(filepath.Join(wd, "../../"))
	pkgConfig := getConfig(root, config.IndexOpts{})
	pkgConfig.Tests = false

	pkgs, err := packages.Load(pkgConfig, "fmt")
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs) != 1 {
		t.Fatalf("Too many packages: %s", pkgs)
	}

	fmtPkg := pkgs[0]

	if !isStandardLib(fmtPkg) {
		t.Fatal("Package was not a builtin package: pre ensure")
	}

	normalizePackage(&config.IndexOpts{}, fmtPkg)

	if fmtPkg.Module.Path != "github.com/golang/go/src" {
		t.Fatalf("Expected normalized module path %q, got %q",
			"github.com/golang/go/src", fmtPkg.Module.Path)
	}
}

type normalizeTestCase struct {
	Raw        string
	Normalized string
}

func TestNormalizePackageSiblingModule(t *testing.T) {
	cases := []struct {
		name            string
		moduleVersion   string
		modulePath      string
		pkgModulePath   string
		expectedVersion string
	}{
		{
			name:            "sibling subpaths in same repo",
			moduleVersion:   "abc123",
			modulePath:      "github.com/scip-code/scip-go/internal/loader",
			pkgModulePath:   "github.com/scip-code/scip-go/internal/config",
			expectedVersion: "abc123",
		},
		{
			name:            "root module and submodule",
			moduleVersion:   "abc123",
			modulePath:      "github.com/scip-code/scip-go",
			pkgModulePath:   "github.com/scip-code/scip-go/submodule",
			expectedVersion: "abc123",
		},
		{
			name:            "different repos",
			moduleVersion:   "abc123",
			modulePath:      "github.com/scip-code/scip-go",
			pkgModulePath:   "github.com/sourcegraph/sourcegraph",
			expectedVersion: ".",
		},
		{
			name:            "sibling module with empty module version",
			moduleVersion:   ".",
			modulePath:      "github.com/scip-code/scip-go/module-a",
			pkgModulePath:   "github.com/scip-code/scip-go/module-b",
			expectedVersion: ".",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			opts := &config.IndexOpts{
				ModulePath:    tc.modulePath,
				ModuleVersion: tc.moduleVersion,
			}
			pkg := &packages.Package{
				PkgPath: tc.pkgModulePath + "/pkg",
				Module: &packages.Module{
					Path:    tc.pkgModulePath,
					Version: "",
				},
			}
			normalizePackage(opts, pkg)
			if pkg.Module.Version != tc.expectedVersion {
				t.Errorf("want %q, got %q", tc.expectedVersion, pkg.Module.Version)
			}
		})
	}
}

func TestNormalizePackageModuleVersion(t *testing.T) {
	cases := []normalizeTestCase{
		{
			Raw:        "v0.0.0-20180920160851-f15b22f93c73",
			Normalized: "f15b22f93c73",
		},
		{
			Raw:        "v0.3.1-0.20230414160720-beea233bdc0b",
			Normalized: "beea233bdc0b",
		},
		{
			Raw:        "v2.0.0-20180818164646-67afb5ed74ec",
			Normalized: "67afb5ed74ec",
		},
		{
			Raw:        "v1.1.1",
			Normalized: "v1.1.1",
		},
		{
			Raw:        "v1.0.0-beta.1",
			Normalized: "v1.0.0-beta.1",
		},
		{
			Raw:        "v0.0.0",
			Normalized: "v0.0.0",
		},
		{
			Raw:        "v2.0.0+incompatible",
			Normalized: "v2.0.0",
		},
		{
			Raw:        "",
			Normalized: ".",
		},
	}

	for _, testCase := range cases {
		pkg := &packages.Package{
			PkgPath: "github.com/fake_name/fake_module/fake_package",
			Module: &packages.Module{
				Path:    "github.com/fake_name/fake_module",
				Version: testCase.Raw,
			},
		}
		normalizePackage(&config.IndexOpts{}, pkg)

		if pkg.Module.Version != testCase.Normalized {
			t.Errorf("want %q, got %q", testCase.Normalized, pkg.Module.Version)
		}
	}
}

func TestPackagePseudoVersion(t *testing.T) {
	pkg := &packages.Package{
		PkgPath: "github.com/alecthomas/template",
		Module: &packages.Module{
			Path:    "github.com/alecthomas/template",
			Version: "v0.0.0-20190718012654-fb15b899a751",
		},
	}

	if !module.IsPseudoVersion(pkg.Module.Version) {
		t.Fatal("Package did not have a pseudo version: pre ensure")
	}

	normalizePackage(&config.IndexOpts{}, pkg)

	if pkg.Module.Version != "fb15b899a751" {
		t.Errorf("Package pseudo-version was not extracted into a sha: want %q, got %q", "fb15b899a751", pkg.Module.Version)
	}
}

func TestPackageWithinModule(t *testing.T) {
	wd, _ := os.Getwd()
	root, _ := filepath.Abs(filepath.Join(wd, "../../"))

	config := getConfig(root, config.IndexOpts{})
	config.Tests = false

	_, err := packages.Load(config, "./...")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStdlibDetection(t *testing.T) {
	std := &packages.Module{Path: "std"}
	userMod := &packages.Module{Path: "github.com/scip-code/scip-go"}

	testCases := []struct {
		pkgPath  string
		module   *packages.Module
		isStdlib bool
	}{
		{"fmt", std, true},
		{"net/http", std, true},
		{"encoding/json", std, true},
		{"fmt", nil, true},
		{"github.com/scip-code/scip-go", userMod, false},
		{"sg/initial", &packages.Module{Path: "sg/initial"}, false},
	}

	for _, tc := range testCases {
		pkg := &packages.Package{PkgPath: tc.pkgPath, Module: tc.module}
		if got := isStandardLib(pkg); got != tc.isStdlib {
			t.Errorf("isStandardLib(%q, module=%v) = %v, want %v",
				tc.pkgPath, tc.module, got, tc.isStdlib)
		}
	}
}
