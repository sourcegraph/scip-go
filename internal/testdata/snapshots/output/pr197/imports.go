  package pr197
//        ^^^^^ definition scip-go/pr197 0.1.test `scip-go/pr197`/
  
  import (
   "context"
//  ^^^^^^^ definition local 0
//  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
   "fmt"
//  ^^^ definition local 1
//  ^^^ reference github.com/golang/go/src go1.22 fmt/
   "net/http"
//  ^^^^^^^^ definition local 2
//  ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
   _ "embed"
//    ^^^^^ reference github.com/golang/go/src go1.22 embed/
  
   . "strings"
//    ^^^^^^^ reference github.com/golang/go/src go1.22 strings/
  
   h "net/http"
// ^ definition local 3
//    ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  )
  
//⌄ enclosing_range_start scip-go/pr197 0.1.test `scip-go/pr197`/UseImports().
  func UseImports() {
//     ^^^^^^^^^^ definition scip-go/pr197 0.1.test `scip-go/pr197`/UseImports().
//     documentation
//     > ```go
//     > func UseImports()
//     > ```
   fmt.Println(context.Background())
// ^^^ reference local 1
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
//             ^^^^^^^ reference local 0
//                     ^^^^^^^^^^ reference github.com/golang/go/src go1.22 context/Background().
   _ = http.StatusOK
//     ^^^^ reference local 3
//          ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/StatusOK.
   _ = h.DefaultClient
//     ^ reference local 3
//       ^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/DefaultClient.
   _ = Contains("hello", "ell")
//     ^^^^^^^^ reference github.com/golang/go/src go1.22 strings/Contains().
  }
//⌃ enclosing_range_end scip-go/pr197 0.1.test `scip-go/pr197`/UseImports().
  
