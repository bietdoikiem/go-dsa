package climbingstairs

// climbStairs returns the number of possible ways to climb the stair with only 1 or 2 steps
// Time complexity: O(n) - Space Complexity: O(n) (Dynamic Programming)
func climbStairs(n int) int {
	memo := make(map[int]int)
	ans := recursiveClimbStairs(0, n, memo)
	return ans
}

// recursiveClimbsStairs is a helper recursive function to explore different steps and return any possible ways
func recursiveClimbStairs(steps int, n int, memo map[int]int) int {
	// Check if the current step milestone has been memoized
	val, ok := memo[n]
	if ok {
		return val
	}
	// Check if it's already reached the end of the stair
	if n == 0 {
		return 1
	}
	// Check for overstep
	if n < 0 {
		return 0
	}
	// Memoization
	memo[n] = recursiveClimbStairs(1, n-1, memo) + recursiveClimbStairs(2, n-2, memo)
	return memo[n]
}
