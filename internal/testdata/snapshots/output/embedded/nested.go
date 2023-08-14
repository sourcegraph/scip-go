  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
//        documentation package embedded
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.21 `net/http`/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src go1.21 `net/http`/Handler# implementation
   http.Handler
// ^^^^ reference github.com/golang/go/src go1.21 `net/http`/
//      ^^^^^^^ definition local 0
//      ^^^^^^^ reference github.com/golang/go/src go1.21 `net/http`/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#Other.
// documentation ```go
  }
  
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/NestedExample().
//     documentation ```go
//                   ^ definition local 1
//                     ^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#
   _ = n.Handler.ServeHTTP
//     ^ reference local 1
//       ^^^^^^^ reference local 0
//               ^^^^^^^^^ reference github.com/golang/go/src go1.21 `net/http`/Handler#ServeHTTP.
   _ = n.ServeHTTP
//     ^ reference local 1
//       ^^^^^^^^^ reference github.com/golang/go/src go1.21 `net/http`/Handler#ServeHTTP.
   _ = n.Other
//     ^ reference local 1
//       ^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#Other.
  }
  
