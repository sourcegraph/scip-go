  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  const MY_THING = 10
//      ^^^^^^^^ definition 0.1.test `sg/initial`/MY_THING.
//      documentation
//      > ```go
//      > const MY_THING untyped int = 10
//      > ```
  const OTHER_THING = MY_THING
//      ^^^^^^^^^^^ definition 0.1.test `sg/initial`/OTHER_THING.
//      documentation
//      > ```go
//      > const OTHER_THING untyped int = 10
//      > ```
//                    ^^^^^^^^ reference 0.1.test `sg/initial`/MY_THING.
  
  func usesMyThing() {
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/usesMyThing().
//     documentation
//     > ```go
//     > func usesMyThing()
//     > ```
   _ = MY_THING
//     ^^^^^^^^ reference 0.1.test `sg/initial`/MY_THING.
  }
  
  var initFunctions = map[string]int{}
//    ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/initFunctions.
//    documentation
//    > ```go
//    > var initFunctions map[string]int
//    > ```
  
