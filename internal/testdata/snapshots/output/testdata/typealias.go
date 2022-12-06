  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  import (
   "sg/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference sg/testdata/internal/secret/
  )
  
  // Type aliased doc
  type SecretBurger = secret.Burger
//     ^^^^^^^^^^^^ definition sg/testdata/SecretBurger#
//     documentation ```go
//     documentation Type aliased doc
//     documentation ```go
//                    ^^^^^^ reference sg/testdata/internal/secret/
//                           ^^^^^^ reference sg/testdata/internal/secret/Burger#
  
  type BadBurger = struct {
//     ^^^^^^^^^ definition sg/testdata/BadBurger#
//     documentation ```go
//     documentation ```go
   Field string
// ^^^^^ definition sg/testdata/BadBurger#Field.
// documentation ```go
  }
  
