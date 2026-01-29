package symbols

import (
	"testing"

	"github.com/sourcegraph/scip/bindings/go/scip"
	"github.com/stretchr/testify/assert"
)

func TestGetSymbolKind_NilObject(t *testing.T) {
	kind := GetSymbolKind(nil)
	assert.Equal(t, scip.SymbolInformation_UnspecifiedKind, kind)
}
