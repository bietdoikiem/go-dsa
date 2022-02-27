package findminimuminsortedarray

// findMin finds the minimum number in rotated sorted array.
// Time Complexity: O(logn)
func findMin(nums []int) int {
	n := len(nums)
	// Check length-1 nums array
	if n == 1 {
		return nums[0]
	}
	// Pointers at two ends (left, right)
	l := 0
	r := n - 1
	var m int // Middle pointer for binary search
	for l < r {
		m = (l + r) / 2
		// Check for "break" point between sorted and unsorted
		if absDiff(l, r) == 1 {
			return min(nums[l], nums[r])
		}
		// Check if the pivot is in the middle of sorted array
		// Early stop for sorted array
		if nums[m] > nums[l] && nums[m] < nums[r] {
			return nums[l]
		}
		// Check if left side is unsorted
		if nums[m] < nums[l] {
			// Shift right to current pivot
			r = m
		}
		// Check if right is unsorted
		if nums[m] > nums[r] {
			// Shift left to current pivot
			l = m
		}
	}
	return -5001
}

// min returns the minimum number between the two
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// absDiff calculate the absolute difference between x and y
func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
