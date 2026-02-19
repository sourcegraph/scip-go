  package testdata
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/
  
  import (
   "github.com/sourcegraph/scip-go/internal/testdata/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata/internal/secret`/
  )
  
  // Type aliased doc
  type SecretBurger = secret.Burger
//     ^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/SecretBurger#
//     documentation
//     > ```go
//     > type SecretBurger = secret.Burger
//     > ```
//     documentation
//     > Type aliased doc
//     documentation
//     > ```go
//     > struct {
//     >     Field int
//     > }
//     > ```
//                    ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata/internal/secret`/
//                           ^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata/internal/secret`/Burger#
  
  type BadBurger = struct {
//     ^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/BadBurger#
//     documentation
//     > ```go
//     > type BadBurger = struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Field string
//     > }
//     > ```
   Field string
// ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/BadBurger#Field.
// documentation
// > ```go
// > struct field Field string
// > ```
  }
  
