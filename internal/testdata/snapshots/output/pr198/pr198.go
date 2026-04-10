  package pr198
//        ^^^^^ definition 0.1.test `sg/pr198`/
//        documentation
//        > package pr198
  
  import "github.com/example/dep"
//        ^^^^^^^^^^^^^^^^^^^^^^ reference github.com/example/dep 0.1.test `github.com/example/dep`/
  
  // Foo is an interface defined downstream of the type that implements it.
  // The dep.T type (from a dependency) implements Foo, and scip-go should
  // emit an external symbol for dep.T with an IsImplementation relationship
  // pointing to Foo.
  type Foo interface {
//     ^^^ definition 0.1.test `sg/pr198`/Foo#
//     documentation
//     > ```go
//     > type Foo interface
//     > ```
//     documentation
//     > Foo is an interface defined downstream of the type that implements it.
//     > The dep.T type (from a dependency) implements Foo, and scip-go should
//     > emit an external symbol for dep.T with an IsImplementation relationship
//     > pointing to Foo.
//     documentation
//     > ```go
//     > interface {
//     >     Bar()
//     > }
//     > ```
   Bar()
// ^^^ definition 0.1.test `sg/pr198`/Foo#Bar.
// documentation
// > ```go
// > func (Foo).Bar()
// > ```
  }
  
//⌄ enclosing_range_start 0.1.test `sg/pr198`/UseFoo().
  func UseFoo(f Foo) {}
//     ^^^^^^ definition 0.1.test `sg/pr198`/UseFoo().
//     documentation
//     > ```go
//     > func UseFoo(f Foo)
//     > ```
//            ^ definition local 0
//              ^^^ reference 0.1.test `sg/pr198`/Foo#
//                    ⌃ enclosing_range_end 0.1.test `sg/pr198`/UseFoo().
  
//⌄ enclosing_range_start 0.1.test `sg/pr198`/Example().
  func Example() {
//     ^^^^^^^ definition 0.1.test `sg/pr198`/Example().
//     documentation
//     > ```go
//     > func Example()
//     > ```
   UseFoo(&dep.T{})
// ^^^^^^ reference 0.1.test `sg/pr198`/UseFoo().
//         ^^^ reference github.com/example/dep 0.1.test `github.com/example/dep`/
//             ^ reference github.com/example/dep 0.1.test `github.com/example/dep`/T#
  }
//⌃ enclosing_range_end 0.1.test `sg/pr198`/Example().
  
