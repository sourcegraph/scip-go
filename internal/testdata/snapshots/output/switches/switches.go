  package switches
//        ^^^^^^^^ definition 0.1.test `sg/switches`/
//        documentation
//        > package switches
  
  // CustomSwitch does the things in a switch
  type CustomSwitch struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/switches`/CustomSwitch#
//     documentation
//     > ```go
//     > type CustomSwitch struct
//     > ```
//     documentation
//     > CustomSwitch does the things in a switch
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
  // Something does some things... and stuff
  func (c *CustomSwitch) Something() bool { return false }
//      ^ definition local 0
//         ^^^^^^^^^^^^ reference 0.1.test `sg/switches`/CustomSwitch#
//                       ^^^^^^^^^ definition 0.1.test `sg/switches`/CustomSwitch#Something().
//                       documentation
//                       > ```go
//                       > func (*CustomSwitch).Something() bool
//                       > ```
//                       documentation
//                       > Something does some things... and stuff
  
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition 0.1.test `sg/switches`/Switch().
//     documentation
//     > ```go
//     > func Switch(interfaceValue interface{}) bool
//     > ```
//            ^^^^^^^^^^^^^^ definition local 1
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 2
//                         ^^^^^^^^^^^^^^ reference local 1
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 2
//         override_documentation
//         > ```go
//         > int
//         > ```
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 2
//          override_documentation
//          > ```go
//          > bool
//          > ```
   case CustomSwitch:
//      ^^^^^^^^^^^^ reference 0.1.test `sg/switches`/CustomSwitch#
    return concreteValue.Something()
//         ^^^^^^^^^^^^^ reference local 2
//         override_documentation
//         > ```go
//         > sg/switches.CustomSwitch
//         > ```
//                       ^^^^^^^^^ reference 0.1.test `sg/switches`/CustomSwitch#Something().
   default:
    return false
   }
  }
  
