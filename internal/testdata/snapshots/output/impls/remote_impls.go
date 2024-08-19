  package impls
//        ^^^^^ reference 0.1.test `sg/impls`/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/
  
  func Something(r http.ResponseWriter) {}
//     ^^^^^^^^^ definition 0.1.test `sg/impls`/Something().
//     documentation
//     > ```go
//     > func Something(r ResponseWriter)
//     > ```
//               ^ definition local 0
//                 ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/ResponseWriter#
  
  type MyWriter struct{}
//     ^^^^^^^^ definition 0.1.test `sg/impls`/MyWriter#
//     documentation
//     > ```go
//     > type MyWriter struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
//     relationship github.com/golang/go/src go1.22 `crypto/tls`/transcriptHash# implementation
//     relationship github.com/golang/go/src go1.22 `internal/bisect`/Writer# implementation
//     relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter# implementation
//     relationship github.com/golang/go/src go1.22 io/Writer# implementation
  
  func (w MyWriter) Header() http.Header        { panic("") }
//      ^ definition local 1
//        ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
//                  ^^^^^^ definition 0.1.test `sg/impls`/MyWriter#Header().
//                  documentation
//                  > ```go
//                  > func (MyWriter).Header() Header
//                  > ```
//                  relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#Header. implementation
//                           ^^^^ reference github.com/golang/go/src go1.22 `net/http`/
//                                ^^^^^^ reference github.com/golang/go/src go1.22 `net/http`/Header#
  func (w MyWriter) Write([]byte) (int, error)  { panic("") }
//      ^ definition local 2
//        ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
//                  ^^^^^ definition 0.1.test `sg/impls`/MyWriter#Write().
//                  documentation
//                  > ```go
//                  > func (MyWriter).Write([]byte) (int, error)
//                  > ```
//                  relationship github.com/golang/go/src go1.22 `crypto/tls`/transcriptHash#Write. implementation
//                  relationship github.com/golang/go/src go1.22 `internal/bisect`/Writer#Write. implementation
//                  relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#Write. implementation
//                  relationship github.com/golang/go/src go1.22 io/Writer#Write. implementation
  func (w MyWriter) WriteHeader(statusCode int) { panic("") }
//      ^ definition local 3
//        ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
//                  ^^^^^^^^^^^ definition 0.1.test `sg/impls`/MyWriter#WriteHeader().
//                  documentation
//                  > ```go
//                  > func (MyWriter).WriteHeader(statusCode int)
//                  > ```
//                  relationship github.com/golang/go/src go1.22 `net/http`/ResponseWriter#WriteHeader. implementation
//                              ^^^^^^^^^^ definition local 4
  
  func Another() {
//     ^^^^^^^ definition 0.1.test `sg/impls`/Another().
//     documentation
//     > ```go
//     > func Another()
//     > ```
   Something(MyWriter{})
// ^^^^^^^^^ reference 0.1.test `sg/impls`/Something().
//           ^^^^^^^^ reference 0.1.test `sg/impls`/MyWriter#
  }
  
