  package initial
//        ^^^^^^^ reference 0.1.test sg/initial/
  
  func UsesBuiltin() int {
//     ^^^^^^^^^^^ definition 0.1.test sg/initial/UsesBuiltin().
//     documentation ```go
   var x int = 5
//     ^ definition local 0
   return x
//        ^ reference local 0
  }
  
