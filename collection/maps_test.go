package collection_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/andrejacobs/go-collection/collection"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestMapUnion(t *testing.T) {
	a := make(map[string]int, 10)
	b := make(map[string]int, 10)

	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("%c", 97+i)
		a[k] = i
		if i%2 == 0 {
			b[k] = i + 10
		}
	}

	for i := 0; i < 5; i++ {
		k := fmt.Sprintf("%c", 65+i)
		b[k] = i + 20
	}

	expected := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4,
		"f": 5, "g": 6, "h": 7, "i": 8, "j": 9,
		"A": 20, "B": 21, "C": 22, "D": 23, "E": 24}

	c := collection.MapUnion(a, b)
	assert.True(t, cmp.Equal(expected, c))
}

func TestMapIntersection(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b := map[string]int{"a": 10, "c": 30, "e": 5, "f": 6}
	expected := map[string]int{"a": 1, "c": 3}
	c := collection.MapIntersection(a, b)
	assert.True(t, cmp.Equal(expected, c))
}

func TestMapPairIntersection(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b := map[string]int{"a": 10, "c": 3, "b": 2, "f": 6, "d": 42}
	expected := map[string]int{"b": 2, "c": 3}
	c := collection.MapPairIntersection(a, b)
	assert.True(t, cmp.Equal(expected, c))
}

func TestMapDifference(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b := map[string]int{"a": 10, "c": 30, "e": 5, "f": 6}
	expected := map[string]int{"b": 2, "d": 4}
	c := collection.MapDifference(a, b)
	assert.True(t, cmp.Equal(expected, c))

	expected = map[string]int{"e": 5, "f": 6}
	c = collection.MapDifference(b, a)
	assert.True(t, cmp.Equal(expected, c))
}

func TestMapSymmetricDifference(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b := map[string]int{"a": 10, "c": 30, "e": 5, "f": 6}
	expected := map[string]int{"b": 2, "d": 4, "e": 5, "f": 6}
	c := collection.MapSymmetricDifference(a, b)
	assert.True(t, cmp.Equal(expected, c))

	c = collection.MapSymmetricDifference(b, a)
	assert.True(t, cmp.Equal(expected, c))
}

func TestMapSortedByValue(t *testing.T) {
	a := map[string]int{"b": 2, "a": 1, "d": 4, "c": 3}
	expected := []collection.KeyValue[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
		{Key: "d", Value: 4},
	}

	sorted := collection.MapSortedByValue(a, collection.Ascending)
	assert.Equal(t, expected, sorted)

	slices.Reverse(expected)
	sorted = collection.MapSortedByValue(a, collection.Descending)
	assert.Equal(t, expected, sorted)
}

func TestMapSortedByAnyValue(t *testing.T) {
	a := map[string]int{"b": 2, "a": 1, "d": 4, "c": 3}
	expected := []collection.KeyValue[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
		{Key: "d", Value: 4},
	}

	sorted := collection.MapSortedByAnyValue(a,
		func(lhs int, rhs int) bool {
			return lhs < rhs
		})
	assert.Equal(t, expected, sorted)
}

func TestMapSortedByKeys(t *testing.T) {
	a := map[string]int{"b": 2, "a": 12, "d": 4, "c": 13}
	expected := []collection.KeyValue[string, int]{
		{Key: "a", Value: 12},
		{Key: "b", Value: 2},
		{Key: "c", Value: 13},
		{Key: "d", Value: 4},
	}

	sorted := collection.MapSortedByKeys(a, collection.Ascending)
	assert.Equal(t, expected, sorted)

	slices.Reverse(expected)
	sorted = collection.MapSortedByKeys(a, collection.Descending)
	assert.Equal(t, expected, sorted)
}
