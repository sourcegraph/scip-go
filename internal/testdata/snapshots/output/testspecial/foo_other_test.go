  package testspecial_test
//        ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testspecial_test`/
//        documentation
//        > package testspecial_test
  
  import (
   "testing"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
   "sg/testspecial"
//  ^^^^^^^^^^^^^^ reference 0.1.test `sg/testspecial`/
  )
  
  func TestFoo_Blackbox(*testing.T) { testspecial.Foo() }
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
//     documentation
//     > ```go
//     > func TestFoo_Blackbox(*T)
//     > ```
//                       ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                               ^ reference github.com/golang/go/src go1.22 testing/T#
//                                    ^^^^^^^^^^^ reference 0.1.test `sg/testspecial`/
//                                                ^^^ reference 0.1.test `sg/testspecial`/Foo().
  
