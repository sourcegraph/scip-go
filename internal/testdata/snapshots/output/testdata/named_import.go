  package testdata
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/
  
  import (
   . "fmt"
//    ^^^ reference github.com/golang/go/src go1.22 fmt/
   h "net/http"
// ^ definition local 0
//    ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  )
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Example().
  func Example() {
//     ^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Example().
//     documentation
//     > ```go
//     > func Example()
//     > ```
   Println(h.CanonicalHeaderKey("accept-encoding"))
// ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//         ^ reference local 0
//           ^^^^^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/CanonicalHeaderKey().
  }
//⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/Example().
  
