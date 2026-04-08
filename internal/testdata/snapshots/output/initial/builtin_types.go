  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/UsesBuiltin().
  func UsesBuiltin() int {
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/UsesBuiltin().
//     documentation
//     > ```go
//     > func UsesBuiltin() int
//     > ```
   var x int = 5
//     ^ definition local 0
   return x
//        ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/UsesBuiltin().
  
