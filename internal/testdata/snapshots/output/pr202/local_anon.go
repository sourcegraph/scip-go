  package pr202
//        ^^^^^ definition 0.1.test `sg/pr202`/
  
  // Anonymous structs in local/function scope.
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/localAnonymousStructs().
  func localAnonymousStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/localAnonymousStructs().
//                           signature_documentation
//                           > func localAnonymousStructs()
   a := struct{ x int }{x: 1}
// ^ definition local 0
//   display_name a
//   signature_documentation
//   > var a struct{x int}
//              ^ definition local 1
//                display_name x
//                signature_documentation
//                > field x int
//                      ^ reference local 1
   b := struct{ x int }{x: 2}
// ^ definition local 2
//   display_name b
//   signature_documentation
//   > var b struct{x int}
//              ^ definition local 3
//                display_name x
//                signature_documentation
//                > field x int
//                      ^ reference local 3
   _ = a.x
//     ^ reference local 0
//       ^ reference local 1
   _ = b.x
//     ^ reference local 2
//       ^ reference local 3
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/localAnonymousStructs().
  
//⌄ enclosing_range_start 0.1.test `sg/pr202`/paramAnonymousStruct().
  func paramAnonymousStruct(p struct{ x int }) int {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr202`/paramAnonymousStruct().
//                          signature_documentation
//                          > func paramAnonymousStruct(p struct{x int}) int
//                          ^ definition local 4
//                            display_name p
//                            signature_documentation
//                            > var p struct{x int}
//                                    ^ definition local 5
//                                      display_name x
//                                      signature_documentation
//                                      > field x int
   return p.x
//        ^ reference local 4
//          ^ reference local 5
  }
//⌃ enclosing_range_end 0.1.test `sg/pr202`/paramAnonymousStruct().
  
