  package server
//        ^^^^^^ reference 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/
  
  import "testing"
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
  func TestStuff(t *testing.T) {
//     ^^^^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
//     documentation
//     > ```go
//     > func TestStuff(t *T)
//     > ```
//               ^ definition local 0
//                  ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
//                          ^ reference github.com/golang/go/src go1.22 testing/T#
   wd := "hello"
// ^^ definition local 1
   repo := "world"
// ^^^^ definition local 2
  
   runCmd(t, wd, "git", "init", repo)
// ^^^^^^ reference 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
//        ^ reference local 0
//           ^^ reference local 1
//                              ^^^^ reference local 2
  }
  
