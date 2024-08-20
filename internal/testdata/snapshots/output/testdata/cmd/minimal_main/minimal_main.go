  package main
//        ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/
//        documentation
//        > package main
  
  type User struct {
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#
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
// documentation
// > ```go
// > struct field Id string
// > ```
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Name.
//     documentation
//     > ```go
//     > struct field Name string
//     > ```
  }
  
  type UserResource struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/UserResource#
//     documentation
//     > ```go
//     > type UserResource struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
  func main() {}
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/main().
//     documentation
//     > ```go
//     > func main()
//     > ```
  
