  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
//⌄ enclosing_range_start 0.1.test `sg/initial`/UsesBuiltin().
  func UsesBuiltin() int {
//     ^^^^^^^^^^^ definition 0.1.test `sg/initial`/UsesBuiltin().
//     kind Function
//     documentation
//     > ```go
//     > func UsesBuiltin() int
//     > ```
   var x int = 5
//     ^ definition local 0
//     kind Variable
   return x
//        ^ reference local 0
  }
//⌃ enclosing_range_end 0.1.test `sg/initial`/UsesBuiltin().
  
