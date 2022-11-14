  package switches
  
  // CustomSwitch does the things in a switch
  type CustomSwitch struct{}
//     ^^^^^^^^^^^^ definition sg/switches/CustomSwitch#
//     documentation CustomSwitch does the things in a switch
  
  // Something does some things... and stuff
  func (c *CustomSwitch) Something() bool { return false }
//      ^ definition local 0
//         ^^^^^^^^^^^^ reference sg/switches/CustomSwitch#
//                       ^^^^^^^^^ definition sg/switches/CustomSwitch#Something().
//                       documentation Something does some things... and stuff
  
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition sg/switches/Switch().
//            ^^^^^^^^^^^^^^ definition local 1
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 2
//                         ^^^^^^^^^^^^^^ reference local 1
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 2
//         override documentation int
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 2
//          override documentation bool
   case CustomSwitch:
//      ^^^^^^^^^^^^ reference sg/switches/CustomSwitch#
    return concreteValue.Something()
//         ^^^^^^^^^^^^^ reference local 2
//         override documentation sg/switches.CustomSwitch
//                       ^^^^^^^^^ reference sg/switches/CustomSwitch#Something().
   default:
    return false
   }
  }
  
