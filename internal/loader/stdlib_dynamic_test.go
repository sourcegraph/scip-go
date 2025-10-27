package loader

import (
	"sync"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestStdlibDetectionNoBazel(t *testing.T) {
	// Add a test-only package to the hardcoded list to verify different behavior
	// between dynamic loading and hardcoded list
	stdPackages["testonly"] = struct{}{}
	t.Cleanup(func() {
		delete(stdPackages, "testonly")
	})

	testCases := []struct {
		pkgPath          string
		isStdlib         bool
		goPackagesDriver string
	}{
		{"fmt", true, ""},
		{"net/http", true, ""},
		{"encoding/json", true, ""},
		{"vendor/golang.org/x/crypto/chacha20", true, ""},
		{"github.com/sourcegraph/scip-go", false, ""},
		{"golang.org/x/tools", false, ""},
		{"vendorapi.com/foo", false, ""},
		{"testonly", false, ""},
		{"testonly", true, "bazel"},
	}

	t.Run("DynamicLoading", func(t *testing.T) {
		for _, tc := range testCases {
			stdlibOnce = sync.Once{}
			t.Setenv("GOPACKAGESDRIVER", tc.goPackagesDriver)

			pkg := &packages.Package{PkgPath: tc.pkgPath}
			if got := IsStandardLib(pkg); got != tc.isStdlib {
				t.Errorf("IsStandardLib(%q) = %v, want %v",
					tc.pkgPath, got, tc.isStdlib)
			}
		}
	})
}
