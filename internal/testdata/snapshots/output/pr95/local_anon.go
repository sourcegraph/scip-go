  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Anonymous structs in local/function scope.
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/localAnonymousStructs().
  func localAnonymousStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/localAnonymousStructs().
//     documentation
//     > ```go
//     > func localAnonymousStructs()
//     > ```
   a := struct{ x int }{x: 1}
// ^ definition local 0
//              ^ definition local 1
//                      ^ reference local 1
   b := struct{ x int }{x: 2}
// ^ definition local 2
//              ^ definition local 3
//                      ^ reference local 3
   _ = a.x
//     ^ reference local 0
//       ^ reference local 1
   _ = b.x
//     ^ reference local 2
//       ^ reference local 3
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/localAnonymousStructs().
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/paramAnonymousStruct().
  func paramAnonymousStruct(p struct{ x int }) int {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/paramAnonymousStruct().
//     documentation
//     > ```go
//     > func paramAnonymousStruct(p struct{x int}) int
//     > ```
//                          ^ definition local 4
//                                    ^ definition local 5
   return p.x
//        ^ reference local 4
//          ^ reference local 5
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/paramAnonymousStruct().
  
