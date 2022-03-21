package longestincreasingsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

+Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

+Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

+Example 4:
Input: nums = [4,10,4,3,8,9]
Output: 3

+Example 5:
Input: [0,1,0,3,2,3]
Output: 4

+Example 6:
Input: [1,3,6,7,9,4,10,5,6]
Output: 6
*/

func Test_lengthOfLIS_1(t *testing.T) {
	assert := assert.New(t)
	ans := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	assert.Equal(4, ans)
}

func Test_lengthOfLIS_2(t *testing.T) {
	assert := assert.New(t)
	ans := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	assert.Equal(4, ans)
}

func Test_lengthOfLIS_3(t *testing.T) {
	assert := assert.New(t)
	ans := lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	assert.Equal(1, ans)
}

func Test_lengthOfLIS_4(t *testing.T) {
	assert := assert.New(t)
	ans := lengthOfLIS([]int{4, 10, 4, 3, 8, 9})
	assert.Equal(3, ans)
}

func Test_lengthOfLIS_5(t *testing.T) {
	assert := assert.New(t)
	ans := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	assert.Equal(4, ans)
}

func Test_lengthOfLIS_6(t *testing.T) {
	assert := assert.New(t)
	ans := lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	assert.Equal(6, ans)
}
