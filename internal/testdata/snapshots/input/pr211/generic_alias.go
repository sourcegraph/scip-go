package pr211

// Map is a generic map type.
type Map[K comparable, V any] struct {
	entries []entry[K, V]
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

// Set is a generic alias that partially instantiates Map.
type Set[K comparable] = Map[K, bool]

// Alias with a tighter constraint.
type OrderedSet[K ~int | ~string] = Set[K]

// Alias of an alias (chained).
type StringSet = Set[string]

// Alias with all params forwarded.
type PairMap[K comparable, V any] = Map[K, V]

func UseAliases() {
	_ = Set[int]{}
	_ = OrderedSet[int]{}
	_ = StringSet{}
	_ = PairMap[string, int]{}
}
