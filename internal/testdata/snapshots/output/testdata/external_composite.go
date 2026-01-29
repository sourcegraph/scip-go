  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/NestedHandler#
//     kind Class
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
//      ^^^^^^^ definition 0.1.test `sg/testdata`/NestedHandler#Handler.
//      kind Field
//      documentation
//      > ```go
//      > struct field Handler net/http.Handler
//      > ```
//      ^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#
   Other int
// ^^^^^ definition 0.1.test `sg/testdata`/NestedHandler#Other.
// kind Field
// documentation
// > ```go
// > struct field Other int
// > ```
  }
  
