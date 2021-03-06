package maximumsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+ Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

+ Example 2:
Input: nums = [1]
Output: 1

+ Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

+ Example 4: All negatives
Input : nums = [-1,-2,-3,-4]
Output: -1

+ Example 5:
Input: nums = [1,2]
Output: 3
*/
func Test_maximumSubArray_1(t *testing.T) {
	assert := assert.New(t)
	// Test
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Println("Max subarray sum:", res)
	assert.Equal(6, res, "Maximum sum of subarray should be 6.")
}

func Test_maximumSubArray_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1}
	res := maxSubArray(nums)
	fmt.Println("Max subarray sum:", res)
	assert.Equal(1, res, "Maximum sum of subarray should be 6.")
}
func Test_maximumSubArray_3(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 4, -1, 7, 8}
	res := maxSubArray(nums)
	fmt.Println("Max subarray sum:", res)
	assert.Equal(23, res, "Maximum sum of subarray should be 6.")
}

func Test_maximumSubArray_4(t *testing.T) {
	nums := []int{-1, -2, -3, -4}
	fmt.Println("Max subarray sum:", maxSubArray(nums))
}

func Test_maximumSubArray_5(t *testing.T) {
	nums := []int{1, 2}
	fmt.Println("Max subarray sum:", maxSubArray(nums))
}
