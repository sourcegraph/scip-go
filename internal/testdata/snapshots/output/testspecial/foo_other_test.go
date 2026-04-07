  package testspecial_test
//        ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testspecial_test`/
  
  import (
   "testing"
//  ^^^^^^^ definition local 0
//  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
   "sg/testspecial"
//  ^^^^^^^^^^^^^^ reference 0.1.test `sg/testspecial`/
//     ^^^^^^^^^^^ definition local 1
  )
  
//⌄ enclosing_range_start 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
  func TestFoo_Blackbox(*testing.T) { testspecial.Foo() }
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
//     documentation
//     > ```go
//     > func TestFoo_Blackbox(*T)
//     > ```
//                       ^^^^^^^ reference local 0
//                               ^ reference github.com/golang/go/src go1.22 testing/T#
//                                    ^^^^^^^^^^^ reference local 1
//                                                ^^^ reference 0.1.test `sg/testspecial`/Foo().
//                                                      ⌃ enclosing_range_end 0.1.test `sg/testspecial_test`/TestFoo_Blackbox().
  
