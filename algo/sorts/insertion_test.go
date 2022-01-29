package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertion_sort(t *testing.T) {
	assert := assert.New(t)
	// Insertion sort
	arr := []int{7, 5, 4, 2}
	InsertionSort(arr)
	Display(arr)
	// Test assertion
	assert.Equal(7, arr[3], "Last element should be 7.")
	assert.Equal(2, arr[0], "First element should be 2.")
}
