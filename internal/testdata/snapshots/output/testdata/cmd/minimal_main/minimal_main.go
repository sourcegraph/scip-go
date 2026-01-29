  package main
//        ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/
//        documentation
//        > package main
  
  type User struct {
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#
//     kind Class
//     documentation
//     > ```go
//     > type User struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Id string
//     >     Name string
//     > }
//     > ```
   Id, Name string
// ^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Id.
// kind Field
// documentation
// > ```go
// > struct field Id string
// > ```
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Name.
//     kind Field
//     documentation
//     > ```go
//     > struct field Name string
//     > ```
  }
  
  type UserResource struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/UserResource#
//     kind Class
//     documentation
//     > ```go
//     > type UserResource struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/cmd/minimal_main`/main().
  func main() {}
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/main().
//     kind Function
//     documentation
//     > ```go
//     > func main()
//     > ```
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/cmd/minimal_main`/main().
  
