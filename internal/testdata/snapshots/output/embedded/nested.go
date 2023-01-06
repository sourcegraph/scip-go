  package embedded
//        ^^^^^^^^ definition sg/embedded/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition sg/embedded/NestedHandler#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src net/http/Handler# implementation
   http.Handler
// ^^^^ reference github.com/golang/go/src net/http/
//      ^^^^^^^ definition sg/embedded/NestedHandler#Handler.
//      documentation ```go
//      ^^^^^^^ reference github.com/golang/go/src net/http/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition sg/embedded/NestedHandler#Other.
// documentation ```go
  }
  
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition sg/embedded/NestedExample().
//     documentation ```go
//                   ^ definition local 0
//                     ^^^^^^^^^^^^^ reference sg/embedded/NestedHandler#
   _ = n.Handler.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^ reference sg/embedded/NestedHandler#Handler.
//               ^^^^^^^^^ reference github.com/golang/go/src net/http/Handler#ServeHTTP.
   _ = n.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^^^ reference github.com/golang/go/src net/http/Handler#ServeHTTP.
   _ = n.Other
//     ^ reference local 0
//       ^^^^^ reference sg/embedded/NestedHandler#Other.
  }
  
