  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import (
   . "fmt"
//  ^^^ reference github.com/golang/go/src v1.19 fmt/
   h "net/http"
//  ^^^^^^^^ reference github.com/golang/go/src v1.19 net/http/
  )
  
  func Example() {
//     ^^^^^^^ definition 0.1.test sg/testdata/Example().
//     documentation ```go
   Println(h.CanonicalHeaderKey("accept-encoding"))
// ^^^^^^^ reference github.com/golang/go/src v1.19 fmt/Println().
//         ^ reference github.com/golang/go/src v1.19 net/http/
//           ^^^^^^^^^^^^^^^^^^ reference github.com/golang/go/src v1.19 net/http/CanonicalHeaderKey().
  }
  
