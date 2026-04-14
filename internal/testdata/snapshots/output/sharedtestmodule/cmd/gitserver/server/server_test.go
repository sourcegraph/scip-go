  package server
//        ^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestExecRequest().
  func TestExecRequest(t *testing.T) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestExecRequest().
//                     signature_documentation
//                     > func TestExecRequest(t *T)
//                     ^ definition local 0
//                       display_name t
//                       signature_documentation
//                       > var t *testing.T
//                        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                                ^ reference github.com/golang/go/src go1.22 testing/T#
   t.Log("hello world")
// ^ reference local 0
//   ^^^ reference github.com/golang/go/src go1.22 testing/common#Log().
  }
//⌃ enclosing_range_end 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestExecRequest().
  
//⌄ enclosing_range_start 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
  func runCmd(t *testing.T, dir string, cmd string, arg ...string) {}
//     ^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
//            signature_documentation
//            > func runCmd(t *T, dir string, cmd string, arg ...string)
//            ^ definition local 1
//              display_name t
//              signature_documentation
//              > var t *testing.T
//               ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                       ^ reference github.com/golang/go/src go1.22 testing/T#
//                          ^^^ definition local 2
//                              display_name dir
//                              signature_documentation
//                              > var dir string
//                                      ^^^ definition local 3
//                                          display_name cmd
//                                          signature_documentation
//                                          > var cmd string
//                                                  ^^^ definition local 4
//                                                      display_name arg
//                                                      signature_documentation
//                                                      > var arg []string
//                                                                  ⌃ enclosing_range_end 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
  
