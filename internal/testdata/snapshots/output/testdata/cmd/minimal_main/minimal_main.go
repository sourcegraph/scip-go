  package main
//        ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/
//             display_name main
//             signature_documentation
//             > package main
  
  type User struct {
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#
//          signature_documentation
//          > type User struct
//          > struct {
//          >     Id string
//          >     Name string
//          > }
   Id, Name string
// ^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Id.
//    signature_documentation
//    > struct field Id string
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/User#Name.
//          signature_documentation
//          > struct field Name string
  }
  
  type UserResource struct{}
//     ^^^^^^^^^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/UserResource#
//                  signature_documentation
//                  > type UserResource struct
//                  > struct{}
  
//⌄ enclosing_range_start 0.1.test `sg/testdata/cmd/minimal_main`/main().
  func main() {}
//     ^^^^ definition 0.1.test `sg/testdata/cmd/minimal_main`/main().
//          signature_documentation
//          > func main()
//             ⌃ enclosing_range_end 0.1.test `sg/testdata/cmd/minimal_main`/main().
  
