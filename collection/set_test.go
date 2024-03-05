package collection_test

import (
	"sort"
	"testing"

	"github.com/andrejacobs/go-collection/collection"
	"github.com/stretchr/testify/assert"
)

func TestSetInsertAndContains(t *testing.T) {
	s := collection.NewSet[int]()

	assert.True(t, s.Insert(5))
	assert.True(t, s.Insert(3))
	assert.True(t, s.Insert(9))
	assert.False(t, s.Insert(5))
	assert.Equal(t, 3, s.Len())

	assert.True(t, s.Contains(5))
	assert.True(t, s.Contains(3))
	assert.True(t, s.Contains(9))
	assert.False(t, s.Contains(42))

	items := s.Items()
	sort.Ints(items)
	assert.Equal(t, []int{3, 5, 9}, items)
}

func TestRemove(t *testing.T) {
	s := collection.NewSet[string]()

	s.Insert("apple")
	s.Insert("pear")
	s.Insert("blueberry")
	assert.Equal(t, 3, s.Len())

	assert.True(t, s.Remove("pear"))
	assert.False(t, s.Contains("pear"))
	assert.Equal(t, 2, s.Len())

	assert.False(t, s.Remove("kiwi"))
}

func TestInsertAndRemoveSlice(t *testing.T) {
	items := []int{5, 9, 3, 42, 3, 42, 5}
	s := collection.NewSet[int]()
	s.InsertSlice(items)
	assert.Equal(t, 4, s.Len())

	s.RemoveSlice([]int{5, 42})
	assert.Equal(t, 2, s.Len())
	assert.False(t, s.Contains(42))
}

func TestNewSetFrom(t *testing.T) {
	items := []int{5, 9, 3, 42, 3, 42, 5}
	s := collection.NewSetFrom(items)
	assert.Equal(t, 4, s.Len())

	items = s.Items()
	sort.Ints(items)
	assert.Equal(t, []int{3, 5, 9, 42}, items)
}

func TestNewSetFromMultipleSlices(t *testing.T) {
	items1 := []int{5, 9, 3, 42, 3, 42, 5}
	items2 := []int{2, 4, 12, 42, 5}
	s := collection.NewSetFrom(items1, items2)
	assert.Equal(t, 7, s.Len())

	items := s.Items()
	sort.Ints(items)
	assert.Equal(t, []int{2, 3, 4, 5, 9, 12, 42}, items)
}

func TestUnion(t *testing.T) {
	a := collection.NewSetFrom([]int{1, 3, 5})
	b := collection.NewSetFrom([]int{2, 4, 6})

	c := a.Union(b)
	assert.Equal(t, 6, c.Len())
}

func TestIntersection(t *testing.T) {
	a := collection.NewSetFrom([]int{1, 3, 5, 42})
	b := collection.NewSetFrom([]int{2, 3, 6, 42})

	c := a.Intersection(b)
	assert.Equal(t, 2, c.Len())
	assert.True(t, c.Contains(3))
	assert.True(t, c.Contains(42))
}

func TestDifference(t *testing.T) {
	a := collection.NewSetFrom([]int{1, 3, 5, 42})
	b := collection.NewSetFrom([]int{2, 3, 6, 42})

	c := a.Difference(b)
	assert.Equal(t, 2, c.Len())
	assert.True(t, c.Contains(1))
	assert.True(t, c.Contains(5))
}

func TestSymmetricDifference(t *testing.T) {
	a := collection.NewSetFrom([]int{1, 3, 5, 42})
	b := collection.NewSetFrom([]int{2, 3, 6, 42})

	c := a.SymmetricDifference(b)
	assert.Equal(t, 4, c.Len())
	assert.True(t, c.Contains(1))
	assert.True(t, c.Contains(2))
	assert.True(t, c.Contains(5))
	assert.True(t, c.Contains(6))
}

func BenchmarkSet(b *testing.B) {
	const iterations = 10000

	bench := func(a collection.Set[int], b collection.Set[int]) {
		for x := 0; x < iterations; x++ {
			a.Insert(x)
			if x%2 == 0 {
				b.Insert(x)
			}
		}

		c := a.Union(b)
		_ = c
		c = a.Intersection(b)
		_ = c
		c = a.Difference(b)
		_ = c
		c = a.SymmetricDifference(b)
		_ = c

		for x := 0; x < iterations/2; x++ {
			_ = a.Contains(x)
		}
	}

	b.Run("NewSet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := collection.NewSet[int]()
			b := collection.NewSet[int]()

			bench(a, b)
		}
	})

	b.Run("NewSetWithCapacity", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := collection.NewSetWithCapacity[int](iterations)
			b := collection.NewSetWithCapacity[int](iterations)

			bench(a, b)
		}
	})
}
