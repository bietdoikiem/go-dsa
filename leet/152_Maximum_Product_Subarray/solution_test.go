package maximumproductsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+ Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

+ Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

/* * Exhaustive/Brute-force approach test * */
func Test_maxProductExhaustive_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 3, -2, 4}
	assert.Equal(6, maxProductExhaustive(nums), "Max product of subarray should be 6.")
}

func Test_maxProductExhaustive_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-2, 0, -1}
	assert.Equal(0, maxProductExhaustive(nums), "Max product of subarray should be 0.")
}
