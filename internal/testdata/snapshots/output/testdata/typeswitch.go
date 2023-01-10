  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition 0.1.test sg/testdata/Switch().
//     documentation ```go
//            ^^^^^^^^^^^^^^ definition local 0
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 1
//                         ^^^^^^^^^^^^^^ reference local 0
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 1
//         override_documentation ```go
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 1
//          override_documentation ```go
   default:
    return false
   }
  }
  
