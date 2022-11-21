  package initial
  
  func UsesLater() {
//     ^^^^^^^^^ definition sg/initial/UsesLater().
//     documentation ```go
   DefinedLater()
// ^^^^^^^^^^^^ reference sg/initial/DefinedLater().
  }
  
  func DefinedLater() {}
//     ^^^^^^^^^^^^ definition sg/initial/DefinedLater().
//     documentation ```go
  
