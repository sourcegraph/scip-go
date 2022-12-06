  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  import (
   . "fmt"
//  ^^^ reference github.com/golang/go/src fmt/
   h "net/http"
//  ^^^^^^^^ reference github.com/golang/go/src net/http/
  )
  
  func Example() {
//     ^^^^^^^ definition sg/testdata/Example().
//     documentation ```go
   Println(h.CanonicalHeaderKey("accept-encoding"))
// ^^^^^^^ reference github.com/golang/go/src fmt/Println().
//         ^ reference github.com/golang/go/src net/http/
//           ^^^^^^^^^^^^^^^^^^ reference github.com/golang/go/src net/http/CanonicalHeaderKey().
  }
  
