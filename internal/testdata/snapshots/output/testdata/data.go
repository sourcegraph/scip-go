  package testdata
//        ^^^^^^^^ definition 0.1.test `sg/testdata`/
//        documentation
//        > package testdata
  
  import (
   "context"
//  ^^^^^^^ reference github.com/golang/go/src go1.22 context/
  
   "sg/testdata/internal/secret"
//  ^^^^^^^^^^^^^^^^^^^^^^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
  )
  
  // TestInterface is an interface used for testing.
  type TestInterface interface {
//     ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestInterface#
//     documentation
//     > ```go
//     > type TestInterface interface
//     > ```
//     documentation
//     > TestInterface is an interface used for testing.
//     documentation
//     > ```go
//     > interface {
//     >     Do(ctx Context, data string) (score int, _ error)
//     > }
//     > ```
   // Do does a test thing.
   Do(ctx context.Context, data string) (score int, _ error)
// ^^ definition 0.1.test `sg/testdata`/TestInterface#Do.
// documentation
// > ```go
// > func (TestInterface).Do(ctx Context, data string) (score int, _ error)
// > ```
//    ^^^ definition local 0
//        ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
//                         ^^^^ definition local 1
//                                       ^^^^^ definition local 2
  }
  
  type (
   // TestStruct is a struct used for testing.
   TestStruct struct {
// ^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#
// documentation
// > ```go
// > type TestStruct struct
// > ```
// documentation
// > ```go
// > struct {
// >     SimpleA int
// >     SimpleB int
// >     SimpleC int
// >     FieldWithTag string "json:\"tag\""
// >     FieldWithAnonymousType struct {
// >         NestedA string
// >         NestedB string
// >         NestedC string
// >     }
// >     EmptyStructField struct{}
// > }
// > ```
    // SimpleA docs
    SimpleA int
//  ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#SimpleA.
//  documentation
//  > ```go
//  > struct field SimpleA int
//  > ```
    // SimpleB docs
    SimpleB int
//  ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#SimpleB.
//  documentation
//  > ```go
//  > struct field SimpleB int
//  > ```
    // SimpleC docs
    SimpleC int
//  ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#SimpleC.
//  documentation
//  > ```go
//  > struct field SimpleC int
//  > ```
  
    FieldWithTag           string `json:"tag"`
//  ^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithTag.
//  documentation
//  > ```go
//  > struct field FieldWithTag string
//  > ```
    FieldWithAnonymousType struct {
//  ^^^^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithAnonymousType.
//  documentation
//  > ```go
//  > struct field FieldWithAnonymousType struct{NestedA string; NestedB string; NestedC string}
//  > ```
     NestedA string
//   ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithAnonymousType.NestedA.
//   documentation
//   > ```go
//   > struct field NestedA string
//   > ```
     NestedB string
//   ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithAnonymousType.NestedB.
//   documentation
//   > ```go
//   > struct field NestedB string
//   > ```
     // NestedC docs
     NestedC string
//   ^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#FieldWithAnonymousType.NestedC.
//   documentation
//   > ```go
//   > struct field NestedC string
//   > ```
    }
  
    EmptyStructField struct{}
//  ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestStruct#EmptyStructField.
//  documentation
//  > ```go
//  > struct field EmptyStructField struct{}
//  > ```
   }
  
   TestEmptyStruct struct{}
// ^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestEmptyStruct#
// documentation
// > ```go
// > type TestEmptyStruct struct
// > ```
// documentation
// > ```go
// > struct{}
// > ```
  )
  
  // Score is just a hardcoded number.
  const Score = uint64(42)
//      ^^^^^ definition 0.1.test `sg/testdata`/Score.
//      documentation
//      > ```go
//      > const Score uint64 = 42
//      > ```
//      documentation
//      > Score is just a hardcoded number.
  const secretScore = secret.SecretScore
//      ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/secretScore.
//      documentation
//      > ```go
//      > const secretScore uint64 = 43
//      > ```
//                    ^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/
//                           ^^^^^^^^^^^ reference 0.1.test `sg/testdata/internal/secret`/SecretScore.
  
  const SomeString = "foobar"
//      ^^^^^^^^^^ definition 0.1.test `sg/testdata`/SomeString.
//      documentation
//      > ```go
//      > const SomeString untyped string = "foobar"
//      > ```
  const LongString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tincidunt viverra aliquam. Phasellus finibus, arcu eu commodo porta, dui quam dictum ante, nec porta enim leo quis felis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Curabitur luctus orci tortor, non condimentum arcu bibendum ut. Proin sit amet vulputate lorem, ut egestas arcu. Curabitur quis sagittis mi. Aenean elit sem, imperdiet ut risus eget, varius varius erat.\nNullam lobortis tortor sed sodales consectetur. Aenean condimentum vehicula elit, eget interdum ante finibus nec. Mauris mollis, nulla eu vehicula rhoncus, eros lectus viverra tellus, ac hendrerit quam massa et felis. Nunc vestibulum diam a facilisis sollicitudin. Aenean nec varius metus. Sed nec diam nibh. Ut erat erat, suscipit et ante eget, tincidunt condimentum orci. Aenean nec facilisis augue, ac sodales ex. Nulla dictum hendrerit tempus. Aliquam fringilla tortor in massa molestie, quis bibendum nulla ullamcorper. Suspendisse congue laoreet elit, vitae consectetur orci facilisis non. Aliquam tempus ultricies sapien, rhoncus tincidunt nisl tincidunt eget. Aliquam nisi ante, rutrum eget viverra imperdiet, congue ut nunc. Donec mollis sed tellus vel placerat. Sed mi ex, fringilla a fermentum a, tincidunt eget lectus.\nPellentesque lacus nibh, accumsan eget feugiat nec, gravida eget urna. Donec quam velit, imperdiet in consequat eget, ultricies eget nunc. Curabitur interdum vel sem et euismod. Donec sed vulputate odio, sit amet bibendum tellus. Integer pellentesque nunc eu turpis cursus, vestibulum sodales ipsum posuere. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Ut at vestibulum sapien. In hac habitasse platea dictumst. Nullam sed lobortis urna, non bibendum ipsum. Sed in sapien quis purus semper fringilla. Integer ut egestas nulla, eu ornare lectus. Maecenas quis sapien condimentum, dignissim urna quis, hendrerit neque. Donec cursus sit amet metus eu mollis.\nSed scelerisque vitae odio non egestas. Cras hendrerit tortor mauris. Aenean quis imperdiet nulla, a viverra purus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent finibus faucibus orci, sed ultrices justo iaculis ut. Ut libero massa, condimentum at elit non, fringilla iaculis quam. Sed sit amet ipsum placerat, tincidunt sem in, efficitur lacus. Curabitur ligula orci, tempus ut magna eget, sodales tristique odio.\nPellentesque in libero ac risus pretium ultrices. In hac habitasse platea dictumst. Curabitur a quam sed orci tempus luctus. Integer commodo nec odio quis consequat. Aenean vitae dapibus augue, nec dictum lectus. Etiam sit amet leo diam. Duis eu ligula venenatis, fermentum lacus vel, interdum odio. Vivamus sit amet libero vitae elit interdum cursus et eu erat. Cras interdum augue sit amet ex aliquet tempor. Praesent dolor nisl, convallis bibendum mauris a, euismod commodo ante. Phasellus non ipsum condimentum, molestie dolor quis, pretium nisi. Mauris augue urna, fermentum ut lacinia a, efficitur vitae odio. Praesent finibus nisl et dolor luctus faucibus. Donec eget lectus sed mi porttitor placerat ac eu odio."
//      ^^^^^^^^^^ definition 0.1.test `sg/testdata`/LongString.
//      documentation
//      > ```go
//      > const LongString untyped string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tincidu...
//      > ```
  const ConstMath = 1 + (2+3)*5
//      ^^^^^^^^^ definition 0.1.test `sg/testdata`/ConstMath.
//      documentation
//      > ```go
//      > const ConstMath untyped int = 26
//      > ```
  
  type StringAlias string
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StringAlias#
//     documentation
//     > ```go
//     > string
//     > ```
  
  const AliasedString StringAlias = "foobar"
//      ^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/AliasedString.
//      documentation
//      > ```go
//      > const AliasedString StringAlias = "foobar"
//      > ```
//                    ^^^^^^^^^^^ reference 0.1.test `sg/testdata`/StringAlias#
  
  // Doer is similar to the test interface (but not the same).
  func (ts *TestStruct) Doer(ctx context.Context, data string) (score int, err error) {
//      ^^ definition local 3
//          ^^^^^^^^^^ reference 0.1.test `sg/testdata`/TestStruct#
//                      ^^^^ definition 0.1.test `sg/testdata`/TestStruct#Doer().
//                      documentation
//                      > ```go
//                      > func (*TestStruct).Doer(ctx Context, data string) (score int, err error)
//                      > ```
//                      documentation
//                      > Doer is similar to the test interface (but not the same).
//                           ^^^ definition local 4
//                               ^^^^^^^ reference github.com/golang/go/src go1.22 context/
//                                       ^^^^^^^ reference github.com/golang/go/src go1.22 context/Context#
//                                                ^^^^ definition local 5
//                                                              ^^^^^ definition local 6
//                                                                         ^^^ definition local 7
   return Score, nil
//        ^^^^^ reference 0.1.test `sg/testdata`/Score.
  }
  
  // StructTagRegression is a struct that caused panic in the wild. Added here to
  // support a regression test.
  //
  // See https://github.com/tal-tech/go-zero/blob/11dd3d75ecceaa3f5772024fb3f26dec1ada8e9c/core/mapping/unmarshaler_test.go#L2272.
  type StructTagRegression struct {
//     ^^^^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/StructTagRegression#
//     documentation
//     > ```go
//     > type StructTagRegression struct
//     > ```
//     documentation
//     > StructTagRegression is a struct that caused panic in the wild. Added here to
//     > support a regression test.
//     > 
//     > See https://github.com/tal-tech/go-zero/blob/11dd3d75ecceaa3f5772024fb3f26dec1ada8e9c/core/mapping/unmarshaler_test.go#L2272.
//     documentation
//     > ```go
//     > struct {
//     >     Value int "key:\",range=[:}\""
//     > }
//     > ```
   Value int `key:",range=[:}"`
// ^^^^^ definition 0.1.test `sg/testdata`/StructTagRegression#Value.
// documentation
// > ```go
// > struct field Value int
// > ```
  }
  
  type TestEqualsStruct = struct {
//     ^^^^^^^^^^^^^^^^ definition 0.1.test `sg/testdata`/TestEqualsStruct#
//     documentation
//     > ```go
//     > type TestEqualsStruct = struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     Value int
//     > }
//     > ```
   Value int
// ^^^^^ definition 0.1.test `sg/testdata`/TestEqualsStruct#Value.
// documentation
// > ```go
// > struct field Value int
// > ```
  }
  
  type ShellStruct struct {
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ShellStruct#
//     documentation
//     > ```go
//     > type ShellStruct struct
//     > ```
//     documentation
//     > ```go
//     > struct {
//     >     InnerStruct
//     > }
//     > ```
   // Ensure this field comes before the definition
   // so that we grab the correct one in our unit
   // tests.
   InnerStruct
// ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/ShellStruct#InnerStruct.
// documentation
// > ```go
// > struct field InnerStruct sg/testdata.InnerStruct
// > ```
// documentation
// > Ensure this field comes before the definition
// > so that we grab the correct one in our unit
// > tests.
// ^^^^^^^^^^^ reference 0.1.test `sg/testdata`/InnerStruct#
  }
  
  type InnerStruct struct{}
//     ^^^^^^^^^^^ definition 0.1.test `sg/testdata`/InnerStruct#
//     documentation
//     > ```go
//     > type InnerStruct struct
//     > ```
//     documentation
//     > ```go
//     > struct{}
//     > ```
  
