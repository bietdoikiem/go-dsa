package containerwithmostwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49

+Example 2:
Input: height = [1,1]
Output: 1
*/

func Test_maxArea_1(t *testing.T) {
	assert := assert.New(t)
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans := maxArea(height)
	assert.Equal(49, ans, "Max area should be 49.")
}

func Test_maxArea_2(t *testing.T) {
	assert := assert.New(t)
	height := []int{1, 1}
	ans := maxArea(height)
	assert.Equal(1, ans, "Max area should be 49.")
}
