  package pr95
//        ^^^^ definition 0.1.test `sg/pr95`/
  
  // Phase 3: Anonymous structs in local/function scope.
  // Local anonymous structs should use local symbols.
  // Two locals with the same anonymous type should still share field symbols.
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/localAnonymousStructs().
  func localAnonymousStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/localAnonymousStructs().
//     documentation
//     > ```go
//     > func localAnonymousStructs()
//     > ```
   a := struct {
// ^ definition local 0
    x int
//  ^ definition local 1
    y string
//  ^ definition local 2
   }{x: 1, y: "one"}
//   ^ reference local 1
//         ^ reference local 2
  
   b := struct {
// ^ definition local 3
    x int
//  ^ definition local 4
    y string
//  ^ definition local 5
   }{x: 2, y: "two"}
//   ^ reference local 4
//         ^ reference local 5
  
   _ = a.x
//     ^ reference local 0
//       ^ reference local 1
   _ = a.y
//     ^ reference local 0
//       ^ reference local 2
   _ = b.x
//     ^ reference local 3
//       ^ reference local 4
   _ = b.y
//     ^ reference local 3
//       ^ reference local 5
  
   // Different type (different fields).
   c := struct {
// ^ definition local 6
    z int
//  ^ definition local 7
   }{z: 3}
//   ^ reference local 7
   _ = c.z
//     ^ reference local 6
//       ^ reference local 7
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/localAnonymousStructs().
  
//⌄ enclosing_range_start 0.1.test `sg/pr95`/paramAnonymousStruct().
  func paramAnonymousStruct(p struct{ x int }) int {
//     ^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/pr95`/paramAnonymousStruct().
//     documentation
//     > ```go
//     > func paramAnonymousStruct(p struct{x int}) int
//     > ```
//                          ^ definition local 8
//                                    ^ definition local 9
   return p.x
//        ^ reference local 8
//          ^ reference local 9
  }
//⌃ enclosing_range_end 0.1.test `sg/pr95`/paramAnonymousStruct().
  
