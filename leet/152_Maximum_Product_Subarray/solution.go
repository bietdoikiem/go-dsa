package maximumproductsubarray

func maxProduct(nums []int) int {
	// TODO: Implement tomorrow!
	return -1
}

// maxProductExhaustive returns the maximum product of subarray using Brute Force/Exhaustive approach
// Time complexity: O(n^2)
// Space complexity: O(1)
func maxProductExhaustive(nums []int) int {
	n := len(nums)
	max := nums[0]
	var prod int
	for i := 0; i < n; i++ {
		prod = nums[i]
		if prod > max {
			max = prod
		}
		for j := i + 1; j < n; j++ {
			prod *= nums[j]
			if prod > max {
				max = prod
			}
		}
	}
	return max
}
