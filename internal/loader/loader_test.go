package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestBuiltinFormat(t *testing.T) {
	wd, _ := os.Getwd()
	config := makeConfig(wd)
	config.Tests = false

	pkgs, err := packages.Load(config, "fmt")
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
	normalizePackage(nil, fmtPkg)

	if !IsStandardLib(fmtPkg) {
		t.Fatal("Package was not a builtin package: post ensure")
	}

	fmt.Println("MODULE", fmtPkg.PkgPath)
	fmt.Println(fmtPkg)
}

func TestPackageWithinModule(t *testing.T) {
	wd, _ := os.Getwd()
	root, _ := filepath.Abs(filepath.Join(wd, "../../"))

	config := makeConfig(root)
	config.Tests = false

	pkgs, err := packages.Load(config, "./...")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pkgs)

}

func TestPentimentoPackage(t *testing.T) {
	// "github.com/efritz/pentimento"
	wd, _ := os.Getwd()
	root, _ := filepath.Abs(filepath.Join(wd, "../../"))

	config := makeConfig(root)
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
