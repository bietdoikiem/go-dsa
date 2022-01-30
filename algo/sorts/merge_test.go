package sorts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	assert := assert.New(t)
	arr := []int{7, 5, 4, 2, 1}
	MergeSort(arr, true)
	fmt.Print("Actual sorted array: ")
	Display(arr)
	// Assertions
	assert.Equal(7, arr[4], "Last element should be 7.")
	assert.Equal(1, arr[0], "First element should be 1.")
	// Test case empty
	emptyArr := []int{}
	emptyResult := MergeSort(emptyArr)
	fmt.Print("Empty sorted array: ")
	Display(emptyResult)
}
