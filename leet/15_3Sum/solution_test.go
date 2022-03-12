package threesum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

+Example 2:
Input: nums = []
Output: []

+Example 3:
Input: nums = [0]
Output: []

+Example 4:
Input: nums = [0,0,0]
Output: [[0,0,0]]

+Example 5:
Input: nums = [0,0,0,0]
Output: [[0,0,0]]

+Example 6:
Input: nums = [-2,0,0,2,2]
Output: [[-2,0,2]]

+Example 7:
Input: nums = [0, -1, 2, -1]
Output: [[-1,-1,0,1]]

+Example 8:
Input: nums = [-2,0,1,1,2]
Output: [[-2,0,2], [-2,1,1]]
*/

func Test_threeSum_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	assert.Equal([][]int{{-1, -1, 2}, {-1, 0, 1}}, ans)
}

func Test_threeSum_4(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 0, 0}
	ans := threeSum(nums)
	assert.Equal([][]int{{0, 0, 0}}, ans)
}

func Test_threeSum_5(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 0, 0, 0}
	ans := threeSum(nums)
	assert.Equal([][]int{{0, 0, 0}}, ans)
}

func Test_threeSum_6(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-2, 0, 0, 2, 2}
	ans := threeSum(nums)
	assert.Equal([][]int{{-2, 0, 2}}, ans)
}

func Test_threeSum_7(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, -1, -1, 1}
	ans := threeSum(nums)
	assert.Equal([][]int{{-1, 0, 1}}, ans)
}

func Test_threeSum_8(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-2, 0, 1, 1, 2}
	ans := threeSum(nums)
	assert.Equal([][]int{{-2, 0, 2}, {-2, 1, 1}}, ans)
}
