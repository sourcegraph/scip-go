  package embedded
//        ^^^^^^^^ definition 0.1.test `sg/embedded`/
//        documentation
//        > package embedded
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  type NestedHandler struct {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#
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
//      ^^^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#Handler.
//      documentation
//      > ```go
//      > struct field Handler net/http.Handler
//      > ```
//      ^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#
  
   // Wow, a great thing for integers
   Other int
// ^^^^^ definition 0.1.test `sg/embedded`/NestedHandler#Other.
// documentation
// > ```go
// > struct field Other int
// > ```
  }
  
  func NestedExample(n NestedHandler) {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/embedded`/NestedExample().
//     documentation
//     > ```go
//     > func NestedExample(n NestedHandler)
//     > ```
//                   ^ definition local 0
//                     ^^^^^^^^^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#
   _ = n.Handler.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#Handler.
//               ^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#ServeHTTP.
   _ = n.ServeHTTP
//     ^ reference local 0
//       ^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Handler#ServeHTTP.
   _ = n.Other
//     ^ reference local 0
//       ^^^^^ reference 0.1.test `sg/embedded`/NestedHandler#Other.
  }
  
