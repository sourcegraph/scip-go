  package initial
//        ^^^^^^^ reference 0.1.test sg/initial/
  
  type (
   // HoverTypeList is a cool struct
   HoverTypeList struct{}
// ^^^^^^^^^^^^^ definition 0.1.test sg/initial/HoverTypeList#
// documentation ```go
// documentation ```go
  )
  
  // This should show up as well
  type HoverType struct{}
//     ^^^^^^^^^ definition 0.1.test sg/initial/HoverType#
//     documentation ```go
//     documentation This should show up as well
//     documentation ```go
  
