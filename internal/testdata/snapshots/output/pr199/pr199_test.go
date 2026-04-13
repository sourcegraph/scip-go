  // Test file documentation for pr199.
  package pr199
//        ^^^^^ definition 0.1.test `sg/pr199`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/pr199`/TestExample().
  func TestExample(t *testing.T) {}
//     ^^^^^^^^^^^ definition 0.1.test `sg/pr199`/TestExample().
//     documentation
//     > ```go
//     > func TestExample(t *T)
//     > ```
//                 ^ definition local 0
//                    ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                            ^ reference github.com/golang/go/src go1.22 testing/T#
//                                ⌃ enclosing_range_end 0.1.test `sg/pr199`/TestExample().
  
