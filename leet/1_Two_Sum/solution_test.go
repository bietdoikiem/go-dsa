package twosum

import (
	"fmt"
	"testing"
)

/*
+ Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

+Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

+Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/
func Test_twoSum_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	ans := twoSum(nums, 9)
	fmt.Println(ans)
}

func Test_twoSum_2(t *testing.T) {
	nums := []int{3, 2, 4}
	ans := twoSum(nums, 6)
	fmt.Println(ans)
}

func Test_twoSum_3(t *testing.T) {
	nums := []int{3, 3}
	ans := twoSum(nums, 6)
	fmt.Println(ans)
}
