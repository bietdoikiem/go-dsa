package longestincreasingsubsequence

// lengthOfLIS returns the length of the longest increasing subsequence
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	ans := helper(nums, 0, -1, dp)
	return ans
}

// helper (brute-force) calculates the maximum length of subsequence in the two cases: take and NOT take current elem
func helper(nums []int, curIdx int, prevIdx int, dp []int) int {
	if curIdx >= len(nums) {
		return 0
	}
	// Check in cache
	if dp[prevIdx+1] != 0 {
		return dp[prevIdx+1]
	}
	// Check for memoization
	notTake := helper(nums, curIdx+1, prevIdx, dp)
	take := 0
	if prevIdx == -1 || nums[curIdx] > nums[prevIdx] {
		take = 1 + helper(nums, curIdx+1, curIdx, dp)
	}
	// Update to memo
	dp[prevIdx+1] = max(take, notTake)
	return dp[prevIdx+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
