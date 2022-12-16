  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition sg/testdata/NestedHandler#
//     documentation ```go
//     documentation ```go
   http.Handler
// ^^^^ reference github.com/golang/go/src net/http/
//      ^^^^^^^ definition sg/testdata/NestedHandler#Handler.
//      documentation ```go
//      ^^^^^^^ reference github.com/golang/go/src net/http/Handler#
   Other int
// ^^^^^ definition sg/testdata/NestedHandler#Other.
// documentation ```go
  }
  
