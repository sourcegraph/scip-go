  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import (
   "sg/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/internal/secret/
  )
  
  // Type aliased doc
  type SecretBurger = secret.Burger
//     ^^^^^^^^^^^^ definition 0.1.test sg/testdata/SecretBurger#
//     documentation ```go
//     documentation Type aliased doc
//     documentation ```go
//                    ^^^^^^ reference 0.1.test sg/testdata/internal/secret/
//                           ^^^^^^ reference 0.1.test sg/testdata/internal/secret/Burger#
  
  type BadBurger = struct {
//     ^^^^^^^^^ definition 0.1.test sg/testdata/BadBurger#
//     documentation ```go
//     documentation ```go
   Field string
// ^^^^^ definition 0.1.test sg/testdata/BadBurger#Field.
// documentation ```go
  }
  
