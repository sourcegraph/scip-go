  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  import (
   "sg/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
  )
  
  // Type aliased doc
  type SecretBurger = secret.Burger
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/SecretBurger#
//     documentation
//     > ```go
//     > type SecretBurger struct
//     > ```
//     documentation
//     > Type aliased doc
//     documentation
//     > ```go
//     > struct {
//     >     Field int
//     > }
//     > ```
//                    ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
//                           ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/Burger#
  
  type BadBurger = struct {
//     ^^^^^^^^^ definition 0.1.test `sg/testdata`/BadBurger#
//     documentation
//     > ```go
//     > type BadBurger struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Field string
//     > }
//     > ```
   Field string
// ^^^^^ definition 0.1.test `sg/testdata`/BadBurger#Field.
// documentation
// > ```go
// > struct field Field string
// > ```
  }
  
