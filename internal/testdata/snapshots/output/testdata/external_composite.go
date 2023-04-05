  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.19 net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test sg/testdata/NestedHandler#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src go1.19 net/http/Handler# implementation
   http.Handler
// ^^^^ reference github.com/golang/go/src go1.19 net/http/
//      ^^^^^^^ definition 0.1.test sg/testdata/NestedHandler#Handler.
//      documentation ```go
//      ^^^^^^^ reference github.com/golang/go/src go1.19 net/http/Handler#
   Other int
// ^^^^^ definition 0.1.test sg/testdata/NestedHandler#Other.
// documentation ```go
  }
  
