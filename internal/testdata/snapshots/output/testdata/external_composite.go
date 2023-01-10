  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import "net/http"
//        ^^^^^^^^ reference v1.19 net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test sg/testdata/NestedHandler#
//     documentation ```go
//     documentation ```go
//     relationship v1.19 net/http/Handler# implementation
   http.Handler
// ^^^^ reference v1.19 net/http/
//      ^^^^^^^ definition 0.1.test sg/testdata/NestedHandler#Handler.
//      documentation ```go
//      ^^^^^^^ reference v1.19 net/http/Handler#
   Other int
// ^^^^^ definition 0.1.test sg/testdata/NestedHandler#Other.
// documentation ```go
  }
  
