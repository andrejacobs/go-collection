// Copyright (c) 2024 Andre Jacobs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
