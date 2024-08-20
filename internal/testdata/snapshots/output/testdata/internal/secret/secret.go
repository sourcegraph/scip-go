  package secret
//        ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
  
  // SecretScore is like score but _secret_.
  const SecretScore = uint64(43)
//      ^^^^^^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/SecretScore.
//      documentation
//      > ```go
//      > const SecretScore uint64 = 43
//      > ```
//      documentation
//      > SecretScore is like score but _secret_.
  
  // Original doc
  type Burger struct {
//     ^^^^^^ definition 0.1.test `sg/testdata/internal/secret`/Burger#
//     documentation
//     > ```go
//     > type Burger struct
//     > ```
//     documentation
//     > Original doc
//     documentation
//     > ```go
//     > struct {
//     >     Field int
//     > }
//     > ```
   Field int
// ^^^^^ definition 0.1.test `sg/testdata/internal/secret`/Burger#Field.
// documentation
// > ```go
// > struct field Field int
// > ```
  }
  
