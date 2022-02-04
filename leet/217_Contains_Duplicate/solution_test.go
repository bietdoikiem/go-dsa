package containsduplicate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+ Example 1:
Input: nums = [1,2,3,1]
Output: true

+ Example 2:
Input: nums = [1,2,3,4]
Output: false

+ Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func Test_containsDuplicate_1(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println("Does the nums array contain duplicate?", result)
	// Assertion
	assert.Equal(true, result, "Array should contain duplicated element.")
}

func Test_containsDuplicate_2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3, 4}
	result := containsDuplicate(nums)
	fmt.Println("Does the nums array contain duplicate?", result)
	// Assertion
	assert.Equal(false, result, "Array should NOT contain duplicated element.")
}

func Test_containsDuplicate_3(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result := containsDuplicate(nums)
	fmt.Println("Does the nums array contain duplicate?", result)
	// Assertion
	assert.Equal(true, result, "Array should NOT contain duplicated element.")
}
