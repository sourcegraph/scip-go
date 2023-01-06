  package testdata
//        ^^^^^^^^ reference sg/testdata/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src net/http/
  
  type implementsWriter struct{}
//     ^^^^^^^^^^^^^^^^ definition sg/testdata/implementsWriter#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src io/Writer# implementation
//     relationship github.com/golang/go/src net/http/ResponseWriter# implementation
  
  func (implementsWriter) Header() http.Header        { panic("Just for how") }
//      ^^^^^^^^^^^^^^^^ reference sg/testdata/implementsWriter#
//                        ^^^^^^ definition sg/testdata/implementsWriter#Header().
//                        documentation ```go
//                        relationship github.com/golang/go/src net/http/ResponseWriter#Header. implementation
//                                 ^^^^ reference github.com/golang/go/src net/http/
//                                      ^^^^^^ reference github.com/golang/go/src net/http/Header#
  func (implementsWriter) Write([]byte) (int, error)  { panic("Just for show") }
//      ^^^^^^^^^^^^^^^^ reference sg/testdata/implementsWriter#
//                        ^^^^^ definition sg/testdata/implementsWriter#Write().
//                        documentation ```go
//                        relationship github.com/golang/go/src io/Writer#Write. implementation
//                        relationship github.com/golang/go/src net/http/ResponseWriter#Write. implementation
  func (implementsWriter) WriteHeader(statusCode int) {}
//      ^^^^^^^^^^^^^^^^ reference sg/testdata/implementsWriter#
//                        ^^^^^^^^^^^ definition sg/testdata/implementsWriter#WriteHeader().
//                        documentation ```go
//                        relationship github.com/golang/go/src net/http/ResponseWriter#WriteHeader. implementation
//                                    ^^^^^^^^^^ definition local 0
  
  func ShowsInSignature(respWriter http.ResponseWriter) {
//     ^^^^^^^^^^^^^^^^ definition sg/testdata/ShowsInSignature().
//     documentation ```go
//                      ^^^^^^^^^^ definition local 1
//                                 ^^^^ reference github.com/golang/go/src net/http/
//                                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src net/http/ResponseWriter#
   respWriter.WriteHeader(1)
// ^^^^^^^^^^ reference local 1
//            ^^^^^^^^^^^ reference github.com/golang/go/src net/http/ResponseWriter#WriteHeader.
  }
  
