  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  // Block doc for the type group.
  type (
   // Alpha is the first type in a group.
   Alpha struct {
// ^^^^^ definition 0.1.test `sg/pr206`/Alpha#
//       signature_documentation
//       > type Alpha struct{ value int }
//       documentation
//       > Block doc for the type group.
    value int
//  ^^^^^ definition 0.1.test `sg/pr206`/Alpha#value.
//        signature_documentation
//        > struct field value int
   }
  
   GammaNoDoc struct {
// ^^^^^^^^^^ definition 0.1.test `sg/pr206`/GammaNoDoc#
//            signature_documentation
//            > type GammaNoDoc struct{ flag bool }
//            documentation
//            > Block doc for the type group.
    flag bool
//  ^^^^ definition 0.1.test `sg/pr206`/GammaNoDoc#flag.
//       signature_documentation
//       > struct field flag bool
   }
  )
  
