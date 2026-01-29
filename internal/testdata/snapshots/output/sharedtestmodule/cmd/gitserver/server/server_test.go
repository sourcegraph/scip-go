  package server
//        ^^^^^^ reference 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestExecRequest().
  func TestExecRequest(t *testing.T) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestExecRequest().
//     kind Function
//     documentation
//     > ```go
//     > func TestExecRequest(t *T)
//     > ```
//                     ^ definition local 0
//                     kind Variable
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
//     kind Function
//     documentation
//     > ```go
//     > func runCmd(t *T, dir string, cmd string, arg ...string)
//     > ```
//            ^ definition local 1
//            kind Variable
//               ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                       ^ reference github.com/golang/go/src go1.22 testing/T#
//                          ^^^ definition local 2
//                          kind Variable
//                                      ^^^ definition local 3
//                                      kind Variable
//                                                  ^^^ definition local 4
//                                                  kind Variable
//                                                                  ⌃ enclosing_range_end 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
  
