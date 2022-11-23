package impls

type I1 interface {
	F1()
}

type I1Clone interface {
	F1()
}

type IfaceOther interface {
	Something()
	Another()
}

type T1 int

func (r T1) F1() {}

type T2 int

func (r T2) F1() {}
func (r T2) F2() {}
