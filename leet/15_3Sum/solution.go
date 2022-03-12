package threesum

import (
	"sort"
)

// threeSum returns list of 3-element subarrays that add up to 0
func threeSum(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	// Skip checking for less than three elements in an array
	if n < 3 {
		return ans
	}
	sort.Ints(nums) // NOTE: in-place quick sort which costs O(n*logn)
	var l int
	var r int
	var s int
	for i := 0; i < n; i++ {
		// Skip duplicated neighbor
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		// Define left, right pointers
		l = i + 1
		r = n - 1
		for l < r {
			// Check for duplicated element in pointers
			// Check for left pointer
			if l > i+1 {
				if nums[l] == nums[l-1] {
					l++
					continue
				}
			}
			// Check for right pointer
			if r < n-1 {
				if nums[r] == nums[r+1] {
					r--
					continue
				}
			}
			s = nums[i] + nums[l] + nums[r]
			if s == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if s > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}
