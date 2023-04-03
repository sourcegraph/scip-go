  package impls
//        ^^^^^ reference 0.1.test sg/impls/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.19 net/http/
  
  func Something(r http.ResponseWriter) {}
//     ^^^^^^^^^ definition 0.1.test sg/impls/Something().
//     documentation ```go
//               ^ definition local 0
//                 ^^^^ reference github.com/golang/go/src go1.19 net/http/
//                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.19 net/http/ResponseWriter#
  
  type MyWriter struct{}
//     ^^^^^^^^ definition 0.1.test sg/impls/MyWriter#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src go1.19 crypto/tls/transcriptHash# implementation
//     relationship github.com/golang/go/src go1.19 io/Writer# implementation
//     relationship github.com/golang/go/src go1.19 net/http/ResponseWriter# implementation
  
  func (w MyWriter) Header() http.Header        { panic("") }
//      ^ definition local 1
//        ^^^^^^^^ reference 0.1.test sg/impls/MyWriter#
//                  ^^^^^^ definition 0.1.test sg/impls/MyWriter#Header().
//                  documentation ```go
//                  relationship github.com/golang/go/src go1.19 net/http/ResponseWriter#Header. implementation
//                           ^^^^ reference github.com/golang/go/src go1.19 net/http/
//                                ^^^^^^ reference github.com/golang/go/src go1.19 net/http/Header#
  func (w MyWriter) Write([]byte) (int, error)  { panic("") }
//      ^ definition local 2
//        ^^^^^^^^ reference 0.1.test sg/impls/MyWriter#
//                  ^^^^^ definition 0.1.test sg/impls/MyWriter#Write().
//                  documentation ```go
//                  relationship github.com/golang/go/src go1.19 crypto/tls/transcriptHash#Write. implementation
//                  relationship github.com/golang/go/src go1.19 io/Writer#Write. implementation
//                  relationship github.com/golang/go/src go1.19 net/http/ResponseWriter#Write. implementation
  func (w MyWriter) WriteHeader(statusCode int) { panic("") }
//      ^ definition local 3
//        ^^^^^^^^ reference 0.1.test sg/impls/MyWriter#
//                  ^^^^^^^^^^^ definition 0.1.test sg/impls/MyWriter#WriteHeader().
//                  documentation ```go
//                  relationship github.com/golang/go/src go1.19 net/http/ResponseWriter#WriteHeader. implementation
//                              ^^^^^^^^^^ definition local 4
  
  func Another() {
//     ^^^^^^^ definition 0.1.test sg/impls/Another().
//     documentation ```go
   Something(MyWriter{})
// ^^^^^^^^^ reference 0.1.test sg/impls/Something().
//           ^^^^^^^^ reference 0.1.test sg/impls/MyWriter#
  }
  
