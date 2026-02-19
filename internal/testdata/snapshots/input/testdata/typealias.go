package testdata

import (
	"github.com/sourcegraph/scip-go/internal/testdata/testdata/internal/secret"
)

// Type aliased doc
type SecretBurger = secret.Burger

type BadBurger = struct {
	Field string
}
