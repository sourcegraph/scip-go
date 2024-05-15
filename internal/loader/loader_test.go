package loader

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/stretchr/testify/require"
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

	if !IsStandardLib(fmtPkg) {
		t.Fatal("Package was not a builtin package: pre ensure")
	}

	// TODO: don't use nil?
	normalizePackage(&config.IndexOpts{}, fmtPkg)

	if !IsStandardLib(fmtPkg) {
		t.Fatal("Package was not a builtin package: post ensure")
	}
}

type normalizeTestCase struct {
	Raw        string
	Normalized string
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
			Normalized: "v2.0.0+incompatible",
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

		require.Equal(t, testCase.Normalized, pkg.Module.Version)
	}
}

func TestPackagePseudoVersion(t *testing.T) {
	wd, _ := os.Getwd()
	root, _ := filepath.Abs(filepath.Join(wd, "../../"))
	pkgConfig := getConfig(root, config.IndexOpts{})
	pkgConfig.Tests = false

	pkgs, err := packages.Load(pkgConfig, "github.com/efritz/pentimento")
	require.Nil(t, err)

	require.Equal(t, 1, len(pkgs), "Too many packages")

	pkg := pkgs[0]

	require.True(t, module.IsPseudoVersion(pkg.Module.Version), "Package did not have a pseudo version: pre ensure")

	normalizePackage(&config.IndexOpts{}, pkg)

	require.Equal(t, "ade47d831101", pkg.Module.Version, "Package pseudo-version was not extracted into a sha: post ensure")
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

func TestPentimentoPackage(t *testing.T) {
	// "github.com/efritz/pentimento"
	wd, _ := os.Getwd()
	root, _ := filepath.Abs(filepath.Join(wd, "../../"))

	config := getConfig(root, config.IndexOpts{})
	config.Tests = false

	// TODO: Could possibly just load this way as well :)
	// packages.Load(config, "github.com/efritz/pentimento")
	pkgs, err := packages.Load(config, "./...")
	if err != nil {
		t.Fatal(err)
	}

	var pentimento *packages.Package
	for _, pkg := range pkgs {
		for _, imported := range pkg.Imports {
			if strings.Contains(imported.Name, "pentimento") {
				pentimento = imported
				break
			}
		}
	}

	if pentimento == nil {
		t.Fatal("Could not find pentimento dep")
	}

	if "pentimento" != pentimento.Name ||
		"github.com/efritz/pentimento" != pentimento.PkgPath ||
		"github.com/efritz/pentimento" != pentimento.Module.Path {

		t.Fatal("Did not match module")
	}
	// Name string = "pentimento"
	// PkgPath string = "github.com/efritz/pentimento"
	// Module:
	//		Path string = "github.com/efritz/pentimento"
	//		Version string = "v0.0.0-20190429011147-ade47d831101"
}
