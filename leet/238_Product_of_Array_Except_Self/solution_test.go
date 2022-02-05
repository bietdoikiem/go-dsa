package productofarrayexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+ Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

+ Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

func Test_productExceptSelf_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3, 4}
	assert.Equal([]int{24, 12, 8, 6}, productExceptSelf(nums), "Array should be the same.")
}

func Test_productExceptSelf_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-1, 1, 0, -3, 3}
	productExceptSelf(nums)
	assert.Equal([]int{0, 0, 9, 0, 0}, productExceptSelf(nums), "Array should be the same.")
}
