  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  // Base provides shared fields.
  type Base struct {
//     ^^^^ definition 0.1.test `sg/pr206`/Base#
//          signature_documentation
//          > type Base struct{ ID int }
//          documentation
//          > Base provides shared fields.
   // ID uniquely identifies the entity.
   ID int
// ^^ definition 0.1.test `sg/pr206`/Base#ID.
//    signature_documentation
//    > struct field ID int
//    documentation
//    > ID uniquely identifies the entity.
  }
  
  type Container struct {
//     ^^^^^^^^^ definition 0.1.test `sg/pr206`/Container#
//               signature_documentation
//               > type Container struct {
//               >     Base
//               >     Extra string
//               > }
   // Base is embedded to inherit shared fields.
   Base
// ^^^^ definition 0.1.test `sg/pr206`/Container#Base.
//      signature_documentation
//      > struct field Base sg/pr206.Base
//      documentation
//      > Base is embedded to inherit shared fields.
// ^^^^ reference 0.1.test `sg/pr206`/Base#
  
   // Extra is a container-specific field.
   Extra string
// ^^^^^ definition 0.1.test `sg/pr206`/Container#Extra.
//       signature_documentation
//       > struct field Extra string
//       documentation
//       > Extra is a container-specific field.
  }
  
