package testdata

import (
	"sg/testdata/internal/secret"
)

// Type aliased doc
type SecretBurger = secret.Burger

type BadBurger = struct {
	Field string
}
