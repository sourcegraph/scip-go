  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  import "fmt"
//        ^^^ reference github.com/golang/go/src go1.22 fmt/
  
  type TypeContainingAnonymousStructs struct {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#
//     documentation ```go
//     documentation ```go
   a, b struct {
// ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
// documentation ```go
//    ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.
//    documentation ```go
    x int
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.x.
//  documentation ```go
    y string
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.y.
//  documentation ```go
   }
   c struct {
// ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
// documentation ```go
    X int
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.X.
//  documentation ```go
    Y string
//  ^ definition 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.Y.
//  documentation ```go
   }
  }
  
  func funcContainingAnonymousStructs() {
//     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/funcContainingAnonymousStructs().
//     documentation ```go
   d := struct {
// ^ definition local 0
    x int
//  ^ definition local 1
    y string
//  ^ definition local 2
   }{
    x: 1,
//  ^ reference local 1
    y: "one",
//  ^ reference local 2
   }
  
   var e struct {
//     ^ definition local 3
    x int
//  ^ definition local 4
    y string
//  ^ definition local 5
   }
  
   e.x = 2
// ^ reference local 3
//   ^ reference local 4
   e.y = "two"
// ^ reference local 3
//   ^ reference local 5
  
   var f TypeContainingAnonymousStructs
//     ^ definition local 6
//       ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#
   f.a.x = 3
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.x.
   f.a.y = "three"
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.y.
   f.b.x = 4
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.x.
   f.b.y = "four"
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.y.
   f.c.X = 5
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.X.
   f.c.Y = "five"
// ^ reference local 6
//   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.Y.
  
   fmt.Printf("> %s, %s\n", d.x, d.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 0
//                            ^ reference local 1
//                               ^ reference local 0
//                                 ^ reference local 2
   fmt.Printf("> %s, %s\n", e.x, e.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 3
//                            ^ reference local 4
//                               ^ reference local 3
//                                 ^ reference local 5
  
   fmt.Printf("> %s, %s\n", f.a.x, f.a.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 6
//                            ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//                              ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.x.
//                                 ^ reference local 6
//                                   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.
//                                     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.y.
   fmt.Printf("> %s, %s\n", f.b.x, f.b.y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 6
//                            ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.
//                              ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.x.
//                                 ^ reference local 6
//                                   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.
//                                     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#a.b.y.
   fmt.Printf("> %s, %s\n", f.c.X, f.c.Y)
// ^^^ reference github.com/golang/go/src go1.22 fmt/
//     ^^^^^^ reference github.com/golang/go/src go1.22 fmt/Printf().
//                          ^ reference local 6
//                            ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//                              ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.X.
//                                 ^ reference local 6
//                                   ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.
//                                     ^ reference 0.1.test `sg/testdata`/TypeContainingAnonymousStructs#c.Y.
  }
  
