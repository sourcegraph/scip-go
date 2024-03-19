  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  import (
   . "fmt"
//    ^^^ reference github.com/golang/go/src go1.22 fmt/
   h "net/http"
// ^ definition local 0
//    ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  )
  
  func Example() {
//     ^^^^^^^ definition 0.1.test `sg/testdata`/Example().
//     documentation ```go
   Println(h.CanonicalHeaderKey("accept-encoding"))
// ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//         ^ reference local 0
//           ^^^^^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/CanonicalHeaderKey().
  }
  
