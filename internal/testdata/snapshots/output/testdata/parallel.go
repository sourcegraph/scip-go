  package testdata
//        ^^^^^^^^ reference 0.1.test sg/testdata/
  
  import (
   "context"
//  ^^^^^^^ reference github.com/golang/go/src go1.19 context/
   "sync"
//  ^^^^ reference github.com/golang/go/src go1.19 sync/
  )
  
  // ParallelizableFunc is a function that can be called concurrently with other instances
  // of this function type.
  type ParallelizableFunc func(ctx context.Context) error
//     ^^^^^^^^^^^^^^^^^^ definition 0.1.test sg/testdata/ParallelizableFunc#
//     documentation ParallelizableFunc is a function that can be called concurrently with other instances
//     documentation ```go
//                             ^^^ definition local 0
//                                 ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//                                         ^^^^^^^ reference github.com/golang/go/src go1.19 context/Context#
  
  // Parallel invokes each of the given parallelizable functions in their own goroutines and
  // returns the first error to occur. This method will block until all goroutines have returned.
  func Parallel(ctx context.Context, fns ...ParallelizableFunc) error {
//     ^^^^^^^^ definition 0.1.test sg/testdata/Parallel().
//     documentation ```go
//     documentation Parallel invokes each of the given parallelizable functions in their own goroutines and
//              ^^^ definition local 1
//                  ^^^^^^^ reference github.com/golang/go/src go1.19 context/
//                          ^^^^^^^ reference github.com/golang/go/src go1.19 context/Context#
//                                   ^^^ definition local 2
//                                          ^^^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/ParallelizableFunc#
   var wg sync.WaitGroup
//     ^^ definition local 3
//        ^^^^ reference github.com/golang/go/src go1.19 sync/
//             ^^^^^^^^^ reference github.com/golang/go/src go1.19 sync/WaitGroup#
   errs := make(chan error, len(fns))
// ^^^^ definition local 4
//                              ^^^ reference local 2
  
   for _, fn := range fns {
//        ^^ definition local 5
//                    ^^^ reference local 2
    wg.Add(1)
//  ^^ reference local 3
//     ^^^ reference github.com/golang/go/src go1.19 sync/WaitGroup#Add().
  
    go func(fn ParallelizableFunc) {
//          ^^ definition local 6
//             ^^^^^^^^^^^^^^^^^^ reference 0.1.test sg/testdata/ParallelizableFunc#
     errs <- fn(ctx)
//   ^^^^ reference local 4
//           ^^ reference local 6
//              ^^^ reference local 1
     wg.Done()
//   ^^ reference local 3
//      ^^^^ reference github.com/golang/go/src go1.19 sync/WaitGroup#Done().
    }(fn)
//    ^^ reference local 5
   }
  
   wg.Wait()
// ^^ reference local 3
//    ^^^^ reference github.com/golang/go/src go1.19 sync/WaitGroup#Wait().
  
   for err := range errs {
//     ^^^ definition local 7
//                  ^^^^ reference local 4
    if err != nil {
//     ^^^ reference local 7
     return err
//          ^^^ reference local 7
    }
   }
  
   return nil
  }
  
