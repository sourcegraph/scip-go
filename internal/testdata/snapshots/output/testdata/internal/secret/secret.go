  package secret
//        ^^^^^^ reference 0.1.test sg/testdata/internal/secret/
  
  // SecretScore is like score but _secret_.
  const SecretScore = uint64(43)
//      ^^^^^^^^^^^ definition 0.1.test sg/testdata/internal/secret/SecretScore.
//      documentation ```go
//      documentation SecretScore is like score but _secret_.
  
  // Original doc
  type Burger struct {
//     ^^^^^^ definition 0.1.test sg/testdata/internal/secret/Burger#
//     documentation ```go
//     documentation Original doc
//     documentation ```go
   Field int
// ^^^^^ definition 0.1.test sg/testdata/internal/secret/Burger#Field.
// documentation ```go
  }
  
