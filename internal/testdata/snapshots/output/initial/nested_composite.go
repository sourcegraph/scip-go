  package initial
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition sg/initial/NestedHandler#
//     documentation ```go
//     documentation ```go
   http.Handler
// ^^^^ reference github.com/golang/go/src net/http/
//      ^^^^^^^ definition sg/initial/NestedHandler#Handler.
//      documentation ```go
//      ^^^^^^^ reference github.com/golang/go/src net/http/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition sg/initial/NestedHandler#Other.
// documentation ```go
  }
  
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition sg/initial/NestedExample().
//     documentation ```go
//                   ^ definition local 0
//                     ^^^^^^^^^^^^^ reference sg/initial/NestedHandler#
   _ = n.Handler.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^ reference sg/initial/NestedHandler#Handler.
//               ^^^^^^^^^ reference github.com/golang/go/src net/http/Handler#ServeHTTP.
   _ = n.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^^^ reference github.com/golang/go/src net/http/Handler#ServeHTTP.
   _ = n.Other
//     ^ reference local 0
//       ^^^^^ reference sg/initial/NestedHandler#Other.
  }
  
