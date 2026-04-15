  package main
//        ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/
//             display_name main
//             signature_documentation
//             > package main
  
  type User struct {
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#
//          display_name User
//          signature_documentation
//          > type User struct {
//          >     Id   string
//          >     Name string
//          > }
   Id, Name string
// ^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Id.
//    display_name Id
//    signature_documentation
//    > struct field Id string
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Name.
//          display_name Name
//          signature_documentation
//          > struct field Name string
  }
  
  type UserResource struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/UserResource#
//                  display_name UserResource
//                  signature_documentation
//                  > type UserResource struct{}
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/cmd/minimal_main`/main().
  func main() {}
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/main().
//          display_name main
//          signature_documentation
//          > func main()
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/cmd/minimal_main`/main().
  
