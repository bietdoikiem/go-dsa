package maximumproductsubarray

// maxProduct returns the largest product of subarray
// Time complexity: O(n)
func maxProduct(nums []int) int {
	n := len(nums)
	// Check edge case: only 1 element
	if n == 1 {
		return nums[0]
	}
	// Find array's largest single element
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	// Storage for current max and min (DP)
	curMax := nums[0]
	curMin := nums[0]
	curMaxProd := 1 // Product of current maximum and current elements
	curMinProd := 1 // Product of current minimum and current elements
	for i := 1; i < n; i++ {
		// Handle edge case '0'
		if nums[i] == 0 {
			curMax = 1
			curMin = 1
			continue
		}
		curMaxProd = curMax * nums[i]
		curMinProd = curMin * nums[i]
		curMax = maxInt(maxInt(curMaxProd, curMinProd), nums[i])
		curMin = minInt(minInt(curMaxProd, curMinProd), nums[i])
		// Check for max
		if curMax > max {
			max = curMax
		}
	}
	return max
}

// maxInt helps finds the maximum value between two integers
func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// minInt helps finds the minimum value between two integers
func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// curMaxProductExhaustive returns the maximum product of subarray using Brute Force/Exhaustive approach
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
