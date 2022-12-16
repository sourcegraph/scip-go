package inlinestruct

type FieldInterface interface {
	SomeMethod() string
}

var MyInline = struct {
	privateField FieldInterface
	PublicField  FieldInterface
}{}

func MyFunc() {
	_ = MyInline.privateField
	_ = MyInline.PublicField
}
