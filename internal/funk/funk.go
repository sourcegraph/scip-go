package funk

func Map[T any, V any](l []T, f func(T) V) []V {
	vals := make([]V, len(l))
	for i, v := range l {
		vals[i] = f(v)
	}

	return vals
}
