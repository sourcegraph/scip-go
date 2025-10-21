package loader

import (
	"os"
	"sync"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestStdlibDetectionNoBazel(t *testing.T) {
	os.Unsetenv("GOPACKAGESDRIVER")

	testCases := []struct {
		pkgPath  string
		isStdlib bool
	}{
		{"fmt", true},
		{"net/http", true},
		{"encoding/json", true},
		{"vendor/golang.org/x/crypto/chacha20", true},
		{"github.com/sourcegraph/scip-go", false},
		{"golang.org/x/tools", false},
		{"vendorapi.com/foo", false},
	}

	t.Run("DynamicLoading", func(t *testing.T) {
		stdlibOnce = sync.Once{}
		stdlibMap = nil

		for _, tc := range testCases {
			pkg := &packages.Package{PkgPath: tc.pkgPath}
			if got := IsStandardLib(pkg); got != tc.isStdlib {
				t.Errorf("IsStandardLib(%q) = %v, want %v",
					tc.pkgPath, got, tc.isStdlib)
			}
		}
	})
}
