  package server
//        ^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/
  
  import "testing"
//        ^^^^^^^ definition local 0
//        ^^^^^^^ reference github.com/golang/go/src go1.22 testing/
  
//⌄ enclosing_range_start 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
  func TestStuff(t *testing.T) {
//     ^^^^^^^^^ definition 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
//     documentation
//     > ```go
//     > func TestStuff(t *T)
//     > ```
//               ^ definition local 1
//                  ^^^^^^^ reference local 0
//                          ^ reference github.com/golang/go/src go1.22 testing/T#
   wd := "hello"
// ^^ definition local 2
   repo := "world"
// ^^^^ definition local 3
  
   runCmd(t, wd, "git", "init", repo)
// ^^^^^^ reference 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/runCmd().
//        ^ reference local 1
//           ^^ reference local 2
//                              ^^^^ reference local 3
  }
//⌃ enclosing_range_end 0.1.test `sg/sharedtestmodule/cmd/gitserver/server`/TestStuff().
  
