package generallyeric

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
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
