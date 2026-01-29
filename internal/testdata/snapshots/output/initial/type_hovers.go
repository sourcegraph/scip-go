  package initial
//        ^^^^^^^ reference 0.1.test `sg/initial`/
  
  type (
   // HoverTypeList is a cool struct
   HoverTypeList struct{}
// ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/HoverTypeList#
// kind Class
// documentation
// > ```go
// > type HoverTypeList struct
// > ```
// documentation
// > ```go
// > struct{}
// > ```
  )
  
  // This should show up as well
  type HoverType struct{}
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/HoverType#
//     kind Class
//     documentation
//     > ```go
//     > type HoverType struct
//     > ```
//     documentation
//     > This should show up as well
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
