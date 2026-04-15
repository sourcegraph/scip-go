package testdata

type I3 interface {
	ScipTestMethod()
}

type EmbeddedI3 interface {
	ScipTestMethod()
}

type TClose struct {
	EmbeddedI3
}
