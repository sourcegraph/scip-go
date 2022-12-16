package symbols

import (
	"fmt"

	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func FromDescriptors(pkg *packages.Package, descriptors ...*scip.Descriptor) string {
	if pkg.Module == nil {
		panic(fmt.Sprintf("Failed to find package for: %+v\n", pkg.PkgPath))
	}

	return scip.VerboseSymbolFormatter.FormatSymbol(&scip.Symbol{
		Scheme: "scip-go",
		Package: &scip.Package{
			Manager: "gomod",
			// TODO: We might not have a dep, so we should handle that
			Name:    pkg.Module.Path,
			Version: pkg.Module.Version,
		},
		Descriptors: descriptors,
	})
}
