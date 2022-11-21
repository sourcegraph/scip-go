  package initial
  
  const MY_THING = 10
//      ^^^^^^^^ definition MY_THING.
//      documentation ```go
  const OTHER_THING = MY_THING
//      ^^^^^^^^^^^ definition OTHER_THING.
//      documentation ```go
//                    ^^^^^^^^ reference MY_THING.
  
  func usesMyThing() {
//     ^^^^^^^^^^^ definition sg/initial/usesMyThing().
//     documentation ```go
   _ = MY_THING
//     ^^^^^^^^ reference MY_THING.
  }
  
