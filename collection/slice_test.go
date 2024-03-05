package collection_test

import (
	"testing"

	"github.com/andrejacobs/go-collection/collection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveAt(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = collection.SliceRemoveAt(s, 2)
	s = collection.SliceRemoveAt(s, 6)
	assert.Equal(t, []int{0, 1, 3, 4, 5, 6, 8, 9}, s)
	require.Panics(t, func() { _ = collection.SliceRemoveAt(s, -1) })
	require.Panics(t, func() { _ = collection.SliceRemoveAt(s, 42) })
}

func TestRemoveAtFast(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = collection.SliceRemoveAtFast(s, 2)
	s = collection.SliceRemoveAtFast(s, 6)
	assert.Equal(t, []int{0, 1, 9, 3, 4, 5, 8, 7}, s)
	require.Panics(t, func() { _ = collection.SliceRemoveAtFast(s, -1) })
	require.Panics(t, func() { _ = collection.SliceRemoveAtFast(s, 42) })
}

func TestSafeRemoveAt(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s, err := collection.SliceSafeRemoveAt(s, 2)
	require.NoError(t, err)
	s, err = collection.SliceSafeRemoveAt(s, 6)
	require.NoError(t, err)
	assert.Equal(t, []int{0, 1, 3, 4, 5, 6, 8, 9}, s)
	_, err = collection.SliceSafeRemoveAt(s, -1)
	assert.Error(t, err)
	_, err = collection.SliceSafeRemoveAt(s, 42)
	assert.Error(t, err)
}

func TestSafeRemoveAtFast(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s, err := collection.SliceSafeRemoveAtFast(s, 2)
	require.NoError(t, err)
	s, err = collection.SliceSafeRemoveAtFast(s, 6)
	require.NoError(t, err)
	assert.Equal(t, []int{0, 1, 9, 3, 4, 5, 8, 7}, s)
	_, err = collection.SliceSafeRemoveAtFast(s, -1)
	assert.Error(t, err)
	_, err = collection.SliceSafeRemoveAtFast(s, 42)
	assert.Error(t, err)
}

func BenchmarkRemoveAt(b *testing.B) {
	const maxItems = 1000
	scenarios := []struct {
		name string
		run  func(items int)
	}{
		{
			name: "RemoveAt",
			run: func(items int) {
				s := make([]int, items)
				for i := 0; i < items/2; i++ {
					s = collection.SliceRemoveAt(s, i)
				}
				_ = s
			},
		},
		{
			name: "RemoveAtFast",
			run: func(items int) {
				s := make([]int, items)
				for i := 0; i < items/2; i++ {
					s = collection.SliceRemoveAtFast(s, i)
				}
				_ = s
			},
		},
		{
			name: "SafeRemoveAt",
			run: func(items int) {
				s := make([]int, items)
				for i := 0; i < items/2; i++ {
					s, _ = collection.SliceSafeRemoveAt(s, i)
				}
				_ = s
			},
		},
		{
			name: "RemoveAtFast",
			run: func(items int) {
				s := make([]int, items)
				for i := 0; i < items/2; i++ {
					s, _ = collection.SliceSafeRemoveAtFast(s, i)
				}
				_ = s
			},
		},
	}

	for _, s := range scenarios {
		b.Run(s.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.run(maxItems)
			}
		})
	}
}

func TestSliceFind(t *testing.T) {
	s := []int{42, 7, 10, 2}
	assert.Equal(t, 0, collection.SliceFind(s, 42))
	assert.Equal(t, 1, collection.SliceFind(s, 7))
	assert.Equal(t, 2, collection.SliceFind(s, 10))
	assert.Equal(t, 3, collection.SliceFind(s, 2))
	assert.Equal(t, -1, collection.SliceFind(s, 100))
}

func TestSliceContains(t *testing.T) {
	s := []int{42, 7, 10, 2}
	assert.True(t, collection.SliceContains(s, 42))
	assert.True(t, collection.SliceContains(s, 7))
	assert.False(t, collection.SliceContains(s, 100))
}
