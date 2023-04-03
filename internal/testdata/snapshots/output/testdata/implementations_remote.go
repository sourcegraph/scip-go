  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import "net/http"
//        ^^^^^^^^ reference github.com/golang/go/src go1.19 net/http/
  
  type implementsWriter struct{}
//     ^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/implementsWriter#
//     documentation ```go
//     documentation ```go
//     relationship github.com/golang/go/src go1.19 crypto/tls/transcriptHash# implementation
//     relationship github.com/golang/go/src go1.19 io/Writer# implementation
//     relationship github.com/golang/go/src go1.19 net/http/ResponseWriter# implementation
  
  func (implementsWriter) Header() http.Header        { panic("Just for how") }
//      ^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/implementsWriter#
//                        ^^^^^^ definition 0.1.test sg/testdata/implementsWriter#Header().
//                        documentation ```go
//                        relationship github.com/golang/go/src go1.19 net/http/ResponseWriter#Header. implementation
//                                 ^^^^ reference github.com/golang/go/src go1.19 net/http/
//                                      ^^^^^^ reference github.com/golang/go/src go1.19 net/http/Header#
  func (implementsWriter) Write([]byte) (int, error)  { panic("Just for show") }
//      ^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/implementsWriter#
//                        ^^^^^ definition 0.1.test sg/testdata/implementsWriter#Write().
//                        documentation ```go
//                        relationship github.com/golang/go/src go1.19 crypto/tls/transcriptHash#Write. implementation
//                        relationship github.com/golang/go/src go1.19 io/Writer#Write. implementation
//                        relationship github.com/golang/go/src go1.19 net/http/ResponseWriter#Write. implementation
  func (implementsWriter) WriteHeader(statusCode int) {}
//      ^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/implementsWriter#
//                        ^^^^^^^^^^^ definition 0.1.test sg/testdata/implementsWriter#WriteHeader().
//                        documentation ```go
//                        relationship github.com/golang/go/src go1.19 net/http/ResponseWriter#WriteHeader. implementation
//                                    ^^^^^^^^^^ definition local 0
  
  func ShowsInSignature(respWriter http.ResponseWriter) {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/ShowsInSignature().
//     documentation ```go
//                      ^^^^^^^^^^ definition local 1
//                                 ^^^^ reference github.com/golang/go/src go1.19 net/http/
//                                      ^^^^^^^^^^^^^^ reference github.com/golang/go/src go1.19 net/http/ResponseWriter#
   respWriter.WriteHeader(1)
// ^^^^^^^^^^ reference local 1
//            ^^^^^^^^^^^ reference github.com/golang/go/src go1.19 net/http/ResponseWriter#WriteHeader.
  }
  
