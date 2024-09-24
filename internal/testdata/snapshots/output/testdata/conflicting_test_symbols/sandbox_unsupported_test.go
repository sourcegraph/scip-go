  //go:build !linux && !windows && !freebsd
  // +build !linux,!windows,!freebsd
  
  package osl
//        ^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/
//        documentation
//        > package osl
  
  import (
   "errors"
//  ^^^^^^ reference github.com/golang/go/src go1.22 errors/
   "testing"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  )
  
  var ErrNotImplemented = errors.New("not implemented")
//    ^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
//    documentation
//    > ```go
//    > var ErrNotImplemented error
//    > ```
//                        ^^^^^^ reference github.com/golang/go/src go1.22 errors/
//                               ^^^ reference github.com/golang/go/src go1.22 errors/New().
  
  func newKey(t *testing.T) (string, error) {
//     ^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/newKey().
//     documentation
//     > ```go
//     > func newKey(t *T) (string, error)
//     > ```
//            ^ definition local 0
//               ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                       ^ reference github.com/golang/go/src go1.22 testing/T#
   return "", ErrNotImplemented
//            ^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/conflicting_test_symbols`/ErrNotImplemented.
  }
  
  func verifySandbox(t *testing.T, s string) {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata/conflicting_test_symbols`/verifySandbox().
//     documentation
//     > ```go
//     > func verifySandbox(t *T, s string)
//     > ```
//                   ^ definition local 1
//                      ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                              ^ reference github.com/golang/go/src go1.22 testing/T#
//                                 ^ definition local 2
   return
  }
  
