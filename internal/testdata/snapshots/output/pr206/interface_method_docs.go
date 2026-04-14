  package pr206
//        ^^^^^ definition 0.1.test `sg/pr206`/
  
  // Doer performs actions.
  type Doer interface {
//     ^^^^ definition 0.1.test `sg/pr206`/Doer#
//          signature_documentation
//          > type Doer interface {
//          >     Do() error
//          >     Reset()
//          > }
//          documentation
//          > Doer performs actions.
   // Do performs the action and returns an error if it fails.
   Do() error
// ^^ definition 0.1.test `sg/pr206`/Doer#Do.
//    signature_documentation
//    > func (Doer).Do() error
  
   // Reset clears internal state.
   Reset()
// ^^^^^ definition 0.1.test `sg/pr206`/Doer#Reset.
//       signature_documentation
//       > func (Doer).Reset()
  }
  
