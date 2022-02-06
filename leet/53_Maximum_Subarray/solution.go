package maximumsubarray

// maxSubArray returns the largest sum of a subarray in the given array
// Time complexity: O(n)
// Space complexity: O(1)
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := max
	for i := 1; i < len(nums); i++ {
		// Get rid of previous results if it does not
		// contribute equally like the current element
		if nums[i] > sum && sum < 0 {
			sum = nums[i]
		} else {
			// Else add the current elem to sum
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
