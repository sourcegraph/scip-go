  //go:build !linux
  // +build !linux
  
  // From https://github.com/moby/moby/blob/master/libnetwork/osl/sandbox_unsupported.go
  
  package osl
//        ^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/
//        documentation package osl
  
  import "errors"
//        ^^^^^^ reference github.com/golang/go/src go1.21 errors/
  
  var (
   // ErrNotImplemented is for platforms which don't implement sandbox
   ErrNotImplemented = errors.New("not implemented")
// ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
// documentation ```go
//                     ^^^^^^ reference github.com/golang/go/src go1.21 errors/
//                            ^^^ reference github.com/golang/go/src go1.21 errors/New().
  )
  
  // NewSandbox provides a new sandbox instance created in an os specific way
  // provided a key which uniquely identifies the sandbox
  func NewSandbox(key string, osCreate, isRestore bool) (*string, error) {
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/NewSandbox().
//     documentation ```go
//     documentation NewSandbox provides a new sandbox instance created in an os specific way
//                ^^^ definition local 0
//                            ^^^^^^^^ definition local 1
//                                      ^^^^^^^^^ definition local 2
   return nil, ErrNotImplemented
//             ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
  }
  
  // GenerateKey generates a sandbox key based on the passed
  // container id.
  func GenerateKey(containerID string) string {
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/GenerateKey().
//     documentation ```go
//     documentation GenerateKey generates a sandbox key based on the passed
//                 ^^^^^^^^^^^ definition local 3
   return ""
  }
  
