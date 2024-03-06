package collection

import (
	"cmp"
	"sort"
)

// See https://en.wikipedia.org/wiki/Set_(mathematics) for set theory.

//FUTURE-TODO: Keep an eye out on changes made to golang.org/x/exp/maps. They now have a Keys method etc.

// Return a new map that is the union of a and b that will contain all of the keys
// from both a and b. If b has the same key as a, then the value of a will be used.
func MapUnion[K comparable, V any](a map[K]V, b map[K]V) map[K]V {
	c := make(map[K]V, len(a)+len(b))

	for k, v := range a {
		c[k] = v
	}

	for k, v := range b {
		_, exists := c[k]
		if !exists {
			c[k] = v
		}
	}

	return c
}

// Return a new map that contains only the keys that are present in both sets.
// Only values from a will be used.
func MapIntersection[K comparable, V any](a map[K]V, b map[K]V) map[K]V {
	c := make(map[K]V)

	for k, v := range a {
		_, exists := b[k]
		if exists {
			c[k] = v
		}
	}

	return c
}

// Return a new map that contains only the key-value pairs that are present in both sets.
// The keys and values of both a and b need to match to qualify for the new map.
func MapPairIntersection[K comparable, V comparable](a map[K]V, b map[K]V) map[K]V {
	c := make(map[K]V)

	for k, aVal := range a {
		bVal, exists := b[k]
		if exists && (aVal == bVal) {
			c[k] = aVal
		}
	}

	return c
}

// Return a new map that contains only the items that are present in a but not in b.
func MapDifference[K comparable, V any](a map[K]V, b map[K]V) map[K]V {
	c := make(map[K]V)

	for k, v := range a {
		_, exists := b[k]
		if !exists {
			c[k] = v
		}
	}

	return c
}

// Return a new map that contains only the items that are present in one or the other map but not the items that appear in both maps.
func MapSymmetricDifference[K comparable, V any](a map[K]V, b map[K]V) map[K]V {
	c := make(map[K]V)

	for k, v := range a {
		_, exists := b[k]
		if !exists {
			c[k] = v
		}
	}

	for k, v := range b {
		_, exists := a[k]
		if !exists {
			c[k] = v
		}
	}

	return c
}

// A Key-Value pair
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

// Return a slice of KeyValue pairs by sorting the values from the specified map
// The value type has to be one of the cmp.Ordered constraints (types that implement <)
func MapSortedByValue[K comparable, V cmp.Ordered](m map[K]V, order SortOrder) []KeyValue[K, V] {
	result := make([]KeyValue[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, KeyValue[K, V]{Key: k, Value: v})
	}

	if order {
		sort.Slice(result, func(i int, j int) bool {
			return result[i].Value < result[j].Value
		})
	} else {
		sort.Slice(result, func(i int, j int) bool {
			return result[j].Value < result[i].Value
		})
	}

	return result
}

// Return a slice of KeyValue pairs by sorting the values from the specified map using the less function provided.
func MapSortedByValueFunc[K comparable, V any](m map[K]V,
	less func(lhs V, rhs V) bool) []KeyValue[K, V] {
	result := make([]KeyValue[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, KeyValue[K, V]{Key: k, Value: v})
	}

	sort.Slice(result, func(i int, j int) bool {
		return less(result[i].Value, result[j].Value)
	})

	return result
}

// Return a slice of KeyValue pairs by sorting the keys from the specified map
// The value type has to be one of the cmp.Ordered constraints (types that implement <)
func MapSortedByKeys[K cmp.Ordered, V any](m map[K]V, order SortOrder) []KeyValue[K, V] {
	result := make([]KeyValue[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, KeyValue[K, V]{Key: k, Value: v})
	}

	if order {
		sort.Slice(result, func(i int, j int) bool {
			return result[i].Key < result[j].Key
		})
	} else {
		sort.Slice(result, func(i int, j int) bool {
			return result[j].Key < result[i].Key
		})
	}

	return result
}
