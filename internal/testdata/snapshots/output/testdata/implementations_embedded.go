  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
  
  type I3 interface {
//     ^^ definition 0.1.test `sg/testdata`/I3#
//        signature_documentation
//        > type I3 interface{ ScipTestMethod() }
   ScipTestMethod()
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/I3#ScipTestMethod.
//                signature_documentation
//                > func (I3).ScipTestMethod()
  }
  
  type EmbeddedI3 interface {
//     ^^^^^^^^^^ definition 0.1.test `sg/testdata`/EmbeddedI3#
//                signature_documentation
//                > type EmbeddedI3 interface{ ScipTestMethod() }
   ScipTestMethod()
// ^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/EmbeddedI3#ScipTestMethod.
//                signature_documentation
//                > func (EmbeddedI3).ScipTestMethod()
//                relationship 0.1.test `sg/testdata`/EmbeddedI3#ScipTestMethod. implementation
//                relationship 0.1.test `sg/testdata`/I3#ScipTestMethod. implementation
  }
  
  type TClose struct {
//     ^^^^^^ definition 0.1.test `sg/testdata`/TClose#
//            signature_documentation
//            > type TClose struct{ EmbeddedI3 }
//            relationship 0.1.test `sg/testdata`/EmbeddedI3# implementation
//            relationship 0.1.test `sg/testdata`/I3# implementation
   EmbeddedI3
// ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TClose#EmbeddedI3.
//            signature_documentation
//            > struct field EmbeddedI3 sg/testdata.EmbeddedI3
// ^^^^^^^^^^ reference 0.1.test `sg/testdata`/EmbeddedI3#
  }
  
