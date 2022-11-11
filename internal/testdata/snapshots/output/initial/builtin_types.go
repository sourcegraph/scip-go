  package initial
  
  func UsesBuiltin() int {
//     ^^^^^^^^^^^ definition sg/initial/UsesBuiltin().
   var x int = 5
//     ^ definition local 0
   return x
//        ^ reference local 0
  }
  
