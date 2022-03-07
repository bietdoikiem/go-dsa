package twosumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

+Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

+Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
*/

func Test_twoSum_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 7, 11, 15}
	ans := twoSum(nums, 9)
	assert.Equal(1, ans[0])
	assert.Equal(2, ans[1])
}

func Test_twoSum_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 3, 4}
	ans := twoSum(nums, 6)
	assert.Equal(1, ans[0])
	assert.Equal(3, ans[1])
}

func Test_twoSum_3(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-1, 0}
	ans := twoSum(nums, -1)
	assert.Equal(1, ans[0])
	assert.Equal(2, ans[1])
}
