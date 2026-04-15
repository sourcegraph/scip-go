  package initial
//        ^^^^^^^ definition 0.1.test `sg/initial`/
  
  type (
   // HoverTypeList is a cool struct
   HoverTypeList struct{}
// ^^^^^^^^^^^^^ definition 0.1.test `sg/initial`/HoverTypeList#
//               signature_documentation
//               > type HoverTypeList struct{}
//               documentation
//               > HoverTypeList is a cool struct
  )
  
  // This should show up as well
  type HoverType struct{}
//     ^^^^^^^^^ definition 0.1.test `sg/initial`/HoverType#
//               signature_documentation
//               > type HoverType struct{}
//               documentation
//               > This should show up as well
  
