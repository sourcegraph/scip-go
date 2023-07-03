  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition 0.1.test sg/testdata/Switch().
//     documentation ```go
//            ^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/Switch().(interfaceValue)
//            documentation ```go
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 0
//                         ^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/Switch().(interfaceValue)
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 0
//         override_documentation ```go
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 0
//          override_documentation ```go
   default:
    return false
   }
  }
  
