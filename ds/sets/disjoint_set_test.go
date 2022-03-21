package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_disjointSet_1(t *testing.T) {
	assert := assert.New(t)
	ds := NewDisjointSet[int]([]int{9, 10, 11, 12, 13, 14, 16})
	ds.Union(1, 2)
	ds.Union(2, 3)
	assert.Equal(ds.connectivity[1], ds.connectivity[2])
	assert.Equal(ds.connectivity[2], ds.connectivity[3])
}
