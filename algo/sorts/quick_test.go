package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quick(t *testing.T) {
	assert := assert.New(t)
	arr := []int{7, 5, 4, 2, 1}
	QuickSort(arr)
	Display(arr)
	// Assertions
	assert.Equal(7, arr[4], "Last element should be 7.")
	assert.Equal(1, arr[0], "First element should be 1.")
	// Test assertion
	assert.Equal(1, 1)
}
