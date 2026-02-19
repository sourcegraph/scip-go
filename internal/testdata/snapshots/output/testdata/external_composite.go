  package testdata
//        ^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/NestedHandler#
//     documentation
//     > ```go
//     > type NestedHandler struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Handler
//     >     Other int
//     > }
//     > ```
//     relationship github.com/golang/go/src go1.22 `net/http`/Handler# implementation
   http.Handler
// ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//      ^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/NestedHandler#Handler.
//      documentation
//      > ```go
//      > struct field Handler net/http.Handler
//      > ```
//      ^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#
   Other int
// ^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testdata`/NestedHandler#Other.
// documentation
// > ```go
// > struct field Other int
// > ```
  }
  
