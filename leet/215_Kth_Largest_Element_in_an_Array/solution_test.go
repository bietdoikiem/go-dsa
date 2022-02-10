package kthlargestelementinanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+ Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

+ Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
*/

func Test_findKthLargestBySorting_1(t *testing.T) {
	assert := assert.New(t) // Assertion
	nums := []int{3, 2, 1, 5, 6, 4}
	assert.Equal(5, findKthLargestBySorting(nums, 2), "2th largest element in the array should be 5.")
}

func Test_findKthLargestByHeap_1(t *testing.T) {
	assert := assert.New(t) // Assertion
	nums := []int{3, 2, 1, 5, 6, 4}
	assert.Equal(5, findKthLargestByHeap(nums, 2), "2th largest element in the array should be 5.")
}

func Test_findKthLargestByQuickSelect_1(t *testing.T) {
	assert := assert.New(t) // Assertion
	nums := []int{3, 2, 1, 5, 6, 4}
	assert.Equal(5, findKthLargestByQuickSelect(nums, 2), "2th largest element in the array should be 5.")
}

func Test_findKthLargestByQuickSelect_2(t *testing.T) {
	assert := assert.New(t) // Assertion
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	assert.Equal(4, findKthLargestByQuickSelect(nums, 4), "2th largest element in the array should be 5.")
}
