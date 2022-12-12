package funk

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Keys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Always give back a sorted list
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	return keys
}

func Map[T any, V any](l []T, f func(T) V) []V {
	vals := make([]V, len(l))
	for i, v := range l {
		vals[i] = f(v)
	}

	return vals
}
