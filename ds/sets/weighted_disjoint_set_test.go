package sets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_weightedDisjointSet_1(t *testing.T) {
	assert := assert.New(t)
	wds := NewWeightedDisjointSet([]int{9, 10, 11, 12, 13, 14, 16})
	wds.Union(0, 1)
	wds.Union(1, 2)
	wds.Union(2, 3)
	fmt.Println("Connection:", wds.connectivity)
	assert.Equal(true, wds.Find(0, 1))
	assert.Equal(true, wds.Find(1, 2))
	assert.Equal(true, wds.Find(2, 3))
}
