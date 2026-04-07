  package testdata
//        ^^^^^^^^ reference 0.1.test `sg/testdata`/
  
  import (
   "context"
//  ^^^^^^^ definition local 0
//  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
   "sync"
//  ^^^^ definition local 1
//  ^^^^ reference github.com/golang/go/src go1.22 sync/
  )
  
  // ParallelizableFunc is a function that can be called concurrently with other instances
  // of this function type.
  type ParallelizableFunc func(ctx context.Context) error
//     ^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ParallelizableFunc#
//     documentation
//     > ParallelizableFunc is a function that can be called concurrently with other instances
//     > of this function type.
//     documentation
//     > ```go
//     > func(ctx Context) error
//     > ```
//                             ^^^ definition local 2
//                                 ^^^^^^^ reference local 0
//                                         ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
  
  // Parallel invokes each of the given parallelizable functions in their own goroutines and
  // returns the first error to occur. This method will block until all goroutines have returned.
//⌄ enclosing_range_start 0.1.test `sg/testdata`/Parallel().
  func Parallel(ctx context.Context, fns ...ParallelizableFunc) error {
//     ^^^^^^^^ definition 0.1.test `sg/testdata`/Parallel().
//     documentation
//     > ```go
//     > func Parallel(ctx Context, fns ...ParallelizableFunc) error
//     > ```
//     documentation
//     > Parallel invokes each of the given parallelizable functions in their own goroutines and
//     > returns the first error to occur. This method will block until all goroutines have returned.
//              ^^^ definition local 3
//                  ^^^^^^^ reference local 0
//                          ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
//                                   ^^^ definition local 4
//                                          ^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/ParallelizableFunc#
   var wg sync.WaitGroup
//     ^^ definition local 5
//        ^^^^ reference local 1
//             ^^^^^^^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#
   errs := make(chan error, len(fns))
// ^^^^ definition local 6
//                              ^^^ reference local 4
  
   for _, fn := range fns {
//        ^^ definition local 7
//                    ^^^ reference local 4
    wg.Add(1)
//  ^^ reference local 5
//     ^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#Add().
  
    go func(fn ParallelizableFunc) {
//          ^^ definition local 8
//             ^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata`/ParallelizableFunc#
     errs <- fn(ctx)
//   ^^^^ reference local 6
//           ^^ reference local 8
//              ^^^ reference local 3
     wg.Done()
//   ^^ reference local 5
//      ^^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#Done().
    }(fn)
//    ^^ reference local 7
   }
  
   wg.Wait()
// ^^ reference local 5
//    ^^^^ reference github.com/golang/go/src go1.22 sync/WaitGroup#Wait().
  
   for err := range errs {
//     ^^^ definition local 9
//                  ^^^^ reference local 6
    if err != nil {
//     ^^^ reference local 9
     return err
//          ^^^ reference local 9
    }
   }
  
   return nil
  }
//⌃ enclosing_range_end 0.1.test `sg/testdata`/Parallel().
  
