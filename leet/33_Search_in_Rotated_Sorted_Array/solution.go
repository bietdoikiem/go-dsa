package searchinrotatedsortedarray

// search searches for the index of the target in the rotated sorted array.
// Return -1 if it's not found in the array
func search(nums []int, target int) int {
	n := len(nums)
	// Skip-checking for one-element array
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l := 0
	r := n - 1
	var m int
	for l < r {
		m = (l + r) / 2 // Calculate middle pointer by averaging left and right ptr
		// Check if indexed elements is the desired result
		if nums[m] == target {
			return m
		}
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}

		// Check for rotated array BEFORE the pivot, pivot is LARGER than left, BIGGER than targer,
		// targer is SMALLER than left-most
		if nums[m] > nums[l] && nums[m] > target && target < nums[l] {
			l = m + 1
			continue
		}
		// ..., pivot is SMALLER than left, SMALLER than target, target LARGER than left-most
		if nums[m] < nums[l] && nums[m] < target && target > nums[l] {
			r = m - 1
			continue
		}

		// If array is already sorted
		// Perform Binary Search
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
