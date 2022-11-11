  package initial
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition sg/initial/NestedHandler#
   http.Handler
//      ^^^^^^^ definition sg/initial/NestedHandler#Handler.
//      ^^^^^^^ reference github.com/golang/go/src net/http/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition sg/initial/NestedHandler#Other.
  }
  
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition sg/initial/NestedExample().
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
  
