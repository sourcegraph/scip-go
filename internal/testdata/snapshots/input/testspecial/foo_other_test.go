package testspecial_test

import (
	"testing"

	"github.com/sourcegraph/scip-go/internal/testdata/testspecial"
)

func TestFoo_Blackbox(*testing.T) { testspecial.Foo() }
