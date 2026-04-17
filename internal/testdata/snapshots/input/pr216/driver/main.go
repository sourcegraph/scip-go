// Command driver is a GOPACKAGESDRIVER proxy for snapshot tests.
// It loads packages via go list (with GOPACKAGESDRIVER=off to prevent
// recursion) and re-emits them as a DriverResponse. Because the driver
// wire format has no Module field, this simulates real build-system
// drivers (Bazel, Buck2) and exercises the normalizePackage code path.
package main

import (
	"encoding/json"
	"os"

	"golang.org/x/tools/go/packages"
)

func main() {
	os.WriteFile(os.Args[0]+".sentinel", nil, 0o644)

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedExportFile,
		Env: append(os.Environ(), "GOPACKAGESDRIVER=off"),
	}, os.Args[1:]...)
	if err != nil {
		os.Exit(1)
	}

	var all []*packages.Package
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		all = append(all, pkg)
	})

	var roots []string
	for _, pkg := range pkgs {
		roots = append(roots, pkg.ID)
	}

	json.NewEncoder(os.Stdout).Encode(packages.DriverResponse{
		Roots:    roots,
		Packages: all,
	})
}
