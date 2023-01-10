  package initial
//        ^^^^^^^ reference 0.1.test sg/initial/
  
  func UsesLater() {
//     ^^^^^^^^^ definition 0.1.test sg/initial/UsesLater().
//     documentation ```go
   DefinedLater()
// ^^^^^^^^^^^^ reference 0.1.test sg/initial/DefinedLater().
  }
  
  func DefinedLater() {}
//     ^^^^^^^^^^^^ definition 0.1.test sg/initial/DefinedLater().
//     documentation ```go
  
