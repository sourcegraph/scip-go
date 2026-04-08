  // Package pr197 tests import symbol definitions.
  package pr197
//        ^^^^^ definition scip-go/pr197 0.1.test `scip-go/pr197`/
//        documentation
//        > Package pr197 tests import symbol definitions.
  
  import "fmt"
//        ^^^ definition local 0
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
//⌄ enclosing_range_start scip-go/pr197 0.1.test `scip-go/pr197`/OtherFunc().
  func OtherFunc() {
//     ^^^^^^^^^ definition scip-go/pr197 0.1.test `scip-go/pr197`/OtherFunc().
//     documentation
//     > ```go
//     > func OtherFunc()
//     > ```
   fmt.Println("from other file")
// ^^^ reference local 0
//     ^^^^^^^ reference github.com/golang/go/src go1.22 fmt/Println().
  }
//⌃ enclosing_range_end scip-go/pr197 0.1.test `scip-go/pr197`/OtherFunc().
  
