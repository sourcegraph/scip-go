  package testdata
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Switch().
  func Switch(interfaceValue interface{}) bool {
//     ^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Switch().
//     documentation
//     > ```go
//     > func Switch(interfaceValue interface{}) bool
//     > ```
//            ^^^^^^^^^^^^^^ definition local 0
   switch concreteValue := interfaceValue.(type) {
//        ^^^^^^^^^^^^^ definition local 1
//                         ^^^^^^^^^^^^^^ reference local 0
   case int:
    return concreteValue*3 > 10
//         ^^^^^^^^^^^^^ reference local 1
//         override_documentation
//         > ```go
//         > int
//         > ```
   case bool:
    return !concreteValue
//          ^^^^^^^^^^^^^ reference local 1
//          override_documentation
//          > ```go
//          > bool
//          > ```
   default:
    return false
   }
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Switch().
  
