  package impls
//        ^^^^^ reference sg/impls/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src net/http/
  
  func Something(r http.ResponseWriter) {}
//     ^^^^^^^^^ definition sg/impls/Something().
//     documentation ```go
//               ^ definition local 0
//                 ^^^^ reference github.com/golang/go/src net/http/
//                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src net/http/ResponseWriter#
  
  type MyWriter struct{}
//     ^^^^^^^^ definition sg/impls/MyWriter#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src io/Writer# implementation
//     relationship github.com/golang/go/src net/http/ResponseWriter# implementation
  
  func (w MyWriter) Header() http.Header        { panic("") }
//      ^ definition local 1
//        ^^^^^^^^ reference sg/impls/MyWriter#
//                  ^^^^^^ definition sg/impls/MyWriter#Header().
//                  documentation ```go
//                  relationship github.com/golang/go/src net/http/ResponseWriter#Header. implementation
//                           ^^^^ reference github.com/golang/go/src net/http/
//                                ^^^^^^ reference github.com/golang/go/src net/http/Header#
  func (w MyWriter) Write([]byte) (int, error)  { panic("") }
//      ^ definition local 2
//        ^^^^^^^^ reference sg/impls/MyWriter#
//                  ^^^^^ definition sg/impls/MyWriter#Write().
//                  documentation ```go
//                  relationship github.com/golang/go/src io/Writer#Write. implementation
//                  relationship github.com/golang/go/src net/http/ResponseWriter#Write. implementation
  func (w MyWriter) WriteHeader(statusCode int) { panic("") }
//      ^ definition local 3
//        ^^^^^^^^ reference sg/impls/MyWriter#
//                  ^^^^^^^^^^^ definition sg/impls/MyWriter#WriteHeader().
//                  documentation ```go
//                  relationship github.com/golang/go/src net/http/ResponseWriter#WriteHeader. implementation
//                              ^^^^^^^^^^ definition local 4
  
  func Another() {
//     ^^^^^^^ definition sg/impls/Another().
//     documentation ```go
   Something(MyWriter{})
// ^^^^^^^^^ reference sg/impls/Something().
//           ^^^^^^^^ reference sg/impls/MyWriter#
  }
  
