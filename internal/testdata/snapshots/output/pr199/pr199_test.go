  // Test file documentation for pr199.
  package pr199
//        ^^^^^ definition 0.1.test `sg/pr199`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/pr199`/TestExample().
  func TestExample(t *testing.T) {}
//     ^^^^^^^^^^^ definition 0.1.test `sg/pr199`/TestExample().
//                 signature_documentation
//                 > func TestExample(t *T)
//                 ^ definition local 0
//                   display_name t
//                   signature_documentation
//                   > var t *T
//                    ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                            ^ reference github.com/golang/go/src go1.22 testing/T#
//                                ⌃ enclosing_range_end 0.1.test `sg/pr199`/TestExample().
  
