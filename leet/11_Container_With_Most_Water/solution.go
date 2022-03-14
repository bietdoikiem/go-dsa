package containerwithmostwater

// maxArea calculate the largest area that can contain water (slanting container not allowed)
// Time complexity: O(n)
func maxArea(height []int) int {
	n := len(height)
	var maxArea int
	var curArea int
	l := 0
	r := n - 1
	for l < r {
		// Calculate current Area
		curArea = (r - l) * min(height[l], height[r])
		maxArea = max(maxArea, curArea)
		// Shift the pointers
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxArea
}

// min returns minimum element between the two
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// max returns maximum element between the two
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
