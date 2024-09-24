  package server
//        ^^^^^^ reference 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
  func TestExecRequest(t *testing.T) {
//     ^^^^^^^^^^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestExecRequest().
//     documentation
//     > ```go
//     > func TestExecRequest(t *T)
//     > ```
//                     ^ definition local 0
//                        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                                ^ reference github.com/golang/go/src go1.22 testing/T#
   t.Log("hello world")
// ^ reference local 0
//   ^^^ reference github.com/golang/go/src go1.22 testing/common#Log().
  }
  
  func runCmd(t *testing.T, dir string, cmd string, arg ...string) {}
//     ^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
//     documentation
//     > ```go
//     > func runCmd(t *T, dir string, cmd string, arg ...string)
//     > ```
//            ^ definition local 1
//               ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                       ^ reference github.com/golang/go/src go1.22 testing/T#
//                          ^^^ definition local 2
//                                      ^^^ definition local 3
//                                                  ^^^ definition local 4
  
