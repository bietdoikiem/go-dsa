package searchinrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

+Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

+Example 3:
Input: nums = [1], target = 0
Output: -1
*/

func Test_search_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	ans := search(nums, 0)
	assert.Equal(4, ans)
}

func Test_search_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	ans := search(nums, 3)
	assert.Equal(-1, ans)
}

func Test_search_3(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1}
	ans := search(nums, 0)
	assert.Equal(-1, ans)
}

// Extra cases
func Test_search_4(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 8, 1, 2}
	ans := search(nums, 8)
	assert.Equal(4, ans)
}

// Extra cases
func Test_search_5(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 8, 9, 2}
	ans := search(nums, 9)
	assert.Equal(5, ans)
}

// Extra cases
func Test_search_6(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 1, 2, 3, 4}
	ans := search(nums, 1)
	assert.Equal(1, ans)
}

func Test_search_7(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 8, 3}
	ans := search(nums, 3)
	assert.Equal(5, ans)
}

func Test_search_8(t *testing.T) {
	assert := assert.New(t)
	nums := []int{8, 9, 2, 3, 4}
	ans := search(nums, 9)
	assert.Equal(1, ans)
}
