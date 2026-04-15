  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  type Base struct{}
//     ^^^^ definition 0.1.test `sg/pr206`/Base#
//          signature_documentation
//          > type Base struct{}
  
  type Container struct {
//     ^^^^^^^^^ definition 0.1.test `sg/pr206`/Container#
//               signature_documentation
//               > type Container struct{ Base }
   // Base is embedded to inherit shared fields.
   Base
// ^^^^ definition 0.1.test `sg/pr206`/Container#Base.
//      signature_documentation
//      > struct field Base sg/pr206.Base
//      documentation
//      > Base is embedded to inherit shared fields.
// ^^^^ reference 0.1.test `sg/pr206`/Base#
  }
  
