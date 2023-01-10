  package inlinestruct
//        ^^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/
  
  type Processor[T any] interface {
//     ^^^^^^^^^ definition 0.1.test sg/inlinestruct/Processor#
//     documentation ```go
//     documentation ```go
//               ^ definition local 0
   Process(payload T)
// ^^^^^^^ definition 0.1.test sg/inlinestruct/Processor#Process.
// documentation ```go
//         ^^^^^^^ definition local 1
//                 ^ reference local 0
   ProcessorType() string
// ^^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/Processor#ProcessorType.
// documentation ```go
  }
  
  type Limit int
//     ^^^^^ definition 0.1.test sg/inlinestruct/Limit#
//     documentation ```go
  
  type ProcessImpl struct{}
//     ^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/ProcessImpl#
//     documentation ```go
//     documentation ```go
  
  func (p *ProcessImpl) Process(payload Limit) { panic("not implemented") }
//      ^ definition local 2
//         ^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/ProcessImpl#
//                      ^^^^^^^ definition 0.1.test sg/inlinestruct/ProcessImpl#Process().
//                      documentation ```go
//                              ^^^^^^^ definition local 3
//                                      ^^^^^ reference 0.1.test sg/inlinestruct/Limit#
  func (p *ProcessImpl) ProcessorType() string { panic("not implemented") }
//      ^ definition local 4
//         ^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/ProcessImpl#
//                      ^^^^^^^^^^^^^ definition 0.1.test sg/inlinestruct/ProcessImpl#ProcessorType().
//                      documentation ```go
  
  var _ Processor[Limit] = &ProcessImpl{}
//      ^^^^^^^^^ reference 0.1.test sg/inlinestruct/Processor#
//                ^^^^^ reference 0.1.test sg/inlinestruct/Limit#
//                          ^^^^^^^^^^^ reference 0.1.test sg/inlinestruct/ProcessImpl#
  
