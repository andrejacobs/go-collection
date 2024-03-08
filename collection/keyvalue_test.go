package collection_test

import (
	"testing"

	"github.com/andrejacobs/go-collection/collection"
	"github.com/stretchr/testify/assert"
)

func TestJustValues(t *testing.T) {
	kvs := []collection.KeyValue[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
	}

	values := collection.JustValues(kvs)
	assert.Equal(t, []int{1, 2, 3}, values)
}
