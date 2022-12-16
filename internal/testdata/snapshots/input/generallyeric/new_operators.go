package generallyeric

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

func Double[T Number](value T) T {
	return value * 2
}

type Box[T any] struct {
	Something T
}

type handler[T any] struct {
	Box[T]
	Another string
}
