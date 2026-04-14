  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  // Block doc for the var group.
  var (
   // BlockVar1 is the first var in a block.
   BlockVar1 = "one"
// ^^^^^^^^^ definition 0.1.test `sg/pr206`/BlockVar1.
//           signature_documentation
//           > var BlockVar1 string
//           documentation
//           > Block doc for the var group.
  
   BlockVarNoDoc = "two"
// ^^^^^^^^^^^^^ definition 0.1.test `sg/pr206`/BlockVarNoDoc.
//               signature_documentation
//               > var BlockVarNoDoc string
//               documentation
//               > Block doc for the var group.
  )
  
