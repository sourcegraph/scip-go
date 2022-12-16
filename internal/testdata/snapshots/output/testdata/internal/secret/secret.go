  package secret
//        ^^^^^^ reference sg/testdata/internal/secret/
  
  // SecretScore is like score but _secret_.
  const SecretScore = uint64(43)
//      ^^^^^^^^^^^ definition SecretScore.
//      documentation ```go
//      documentation SecretScore is like score but _secret_.
  
  // Original doc
  type Burger struct {
//     ^^^^^^ definition sg/testdata/internal/secret/Burger#
//     documentation ```go
//     documentation Original doc
//     documentation ```go
   Field int
// ^^^^^ definition sg/testdata/internal/secret/Burger#Field.
// documentation ```go
  }
  
