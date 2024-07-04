package funk

import (
	"slices"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

func Map[T any, V any](l []T, f func(T) V) []V {
	vals := make([]V, 0, len(l))
	for i, v := range l {
		vals[i] = f(v)
	}

	return vals
}
