  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  const MY_THING = 10
//      ^^^^^^^^ definition 0.1.test `sg/initial`/MY_THING.
//      kind Constant
//      documentation
//      > ```go
//      > const MY_THING untyped int = 10
//      > ```
  const OTHER_THING = MY_THING
//      ^^^^^^^^^^^ definition 0.1.test `sg/initial`/OTHER_THING.
//      kind Constant
//      documentation
//      > ```go
//      > const OTHER_THING untyped int = 10
//      > ```
//                    ^^^^^^^^ reference 0.1.test `sg/initial`/MY_THING.
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/usesMyThing().
  func usesMyThing() {
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/usesMyThing().
//     kind Function
//     documentation
//     > ```go
//     > func usesMyThing()
//     > ```
   _ = MY_THING
//     ^^^^^^^^ reference 0.1.test `sg/initial`/MY_THING.
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/usesMyThing().
  
  var initFunctions = map[string]int{}
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/initFunctions.
//    kind Variable
//    documentation
//    > ```go
//    > var initFunctions map[string]int
//    > ```
  
