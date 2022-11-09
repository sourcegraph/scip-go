  package initial
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go net/http/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition sg/initial/NestedHandler#
   http.Handler
// ^^^^ reference net/http/http/
//      ^^^^^^^ definition local 0
//      ^^^^^^^ reference github.com/golang/go net/http/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition sg/initial/NestedHandler#Other.
// documentation Wow, a great thing for integers
//       ^^^ reference builtin/builtin builtin/int#
  }
  
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition sg/initial/NestedExample().
//                   ^ definition local 1
//                     ^^^^^^^^^^^^^ reference sg/initial/NestedHandler#
   _ = n.Handler.ServeHTTP
//     ^ reference local 1
//       ^^^^^^^ reference local 0
//               ^^^^^^^^^ reference github.com/golang/go net/http/ServeHTTP().
   _ = n.ServeHTTP
//     ^ reference local 1
//       ^^^^^^^^^ reference github.com/golang/go net/http/ServeHTTP().
   _ = n.Other
//     ^ reference local 1
//       ^^^^^ reference sg/initial/NestedHandler#Other.
  }
  
