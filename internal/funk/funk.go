package funk

import (
	"cmp"
	"maps"
	"slices"
)

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	return slices.Sorted(maps.Keys(m))
}

func Map[T any, V any](l []T, f func(T) V) []V {
	vals := make([]V, 0, len(l))
	for i, v := range l {
		vals[i] = f(v)
	}

	return vals
}
