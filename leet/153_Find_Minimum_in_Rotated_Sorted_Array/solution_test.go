package findminimuminsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
+Example 1:
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

+Example 2:
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

+Example 3:
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

+Example 4:
Input: nums = [2, 3, 1]
Output: 1
*/

func Test_findMin_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{3, 4, 5, 1, 2}
	assert.Equal(1, findMin(nums), "Minimum should be 1.")
}

func Test_findMin_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(0, findMin(nums), "Minimum should be 0.")
}
func Test_findMin_3(t *testing.T) {
	assert := assert.New(t)
	nums := []int{11, 13, 15, 17}
	assert.Equal(11, findMin(nums), "Minimum should be 11.")
}
func Test_findMin_4(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 3, 1}
	assert.Equal(1, findMin(nums), "Minimum should be 1.")
}
