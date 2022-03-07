package twosumii

// twoSum returns an array containing indices of two elements that add ups to the target
// Time complexity: O(n)
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	ans := make([]int, 2)
	l := 0
	r := n - 1
	var sum int
	for l < r {
		sum = numbers[l] + numbers[r]
		// Check if equal target
		if sum == target {
			ans[0] = l + 1
			ans[1] = r + 1
			break
		}

		if sum > target {
			// Shift right-most
			r--
		} else {
			// Shift left-most
			l++
		}
	}
	return ans
}
