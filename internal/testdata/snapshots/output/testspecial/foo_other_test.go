  package testspecial_test
//        ^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial_test`/
//        documentation
//        > package testspecial_test
  
  import (
   "testing"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
   "github.com/sourcegraph/scip-go/internal/testdata/testspecial"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial`/
  )
  
//⌄ enclosing_range_start 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial_test`/TestFoo_Blackbox().
  func TestFoo_Blackbox(*testing.T) { testspecial.Foo() }
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial_test`/TestFoo_Blackbox().
//     documentation
//     > ```go
//     > func TestFoo_Blackbox(*T)
//     > ```
//                       ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                               ^ reference github.com/golang/go/src go1.22 testing/T#
//                                    ^^^^^^^^^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial`/
//                                                ^^^ reference 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial`/Foo().
//                                                      ⌃ enclosing_range_end 0.1.test `github.com/sourcegraph/scip-go/internal/testdata/testspecial_test`/TestFoo_Blackbox().
  
