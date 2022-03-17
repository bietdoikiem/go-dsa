package coinchange

var upperBound int = 10001

// coinChange returns the minimum number of coins to return from the given amount.
// Return -1 if there are no possible number of coins to return.
func coinChange(coins []int, amount int) int {
	// Check for zero amount case
	if amount == 0 {
		return 0
	}
	// Memo technique
	memo := make([]int, amount+1)
	// Pre-initialized coins that match the amount
	for _, c := range coins {
		memo[c] = 1
	}
	ans := helper(coins, amount, memo)
	if ans == upperBound {
		return -1
	}
	return ans
}

// helper recursively find the minimum number of coins as change to the given amount
func helper(coins []int, amount int, memo []int) int {
	// Check for memoization
	if memo[amount] != 0 {
		return memo[amount]
	}
	// Return as a valid way if it's equal the given amount
	if amount == 0 {
		return 1
	}
	// Check for invalid coin change as it's over the given amount
	if amount < 0 {
		return -1
	}
	minCoins := upperBound
	curCoins := 0
	var amountLeft int
	for i := 0; i < len(coins); i++ {
		amountLeft = amount - coins[i]
		// Check if coin exceeds the total amount
		if amountLeft < 0 {
			continue
		}
		// Check for cached result
		if memo[amountLeft] != 0 {
			curCoins = memo[amountLeft]
		} else {
			// Update amount left memo
			memo[amountLeft] = helper(coins, amountLeft, memo)
			curCoins = memo[amountLeft]
		}
		// Skip check min for case of invalid current coins
		if curCoins == -1 {
			continue
		}
		minCoins = min(minCoins, 1+curCoins)
	}
	// In case there is no valid minimum value
	if minCoins == upperBound {
		return minCoins
	}
	// Update memoized value
	memo[amount] = minCoins
	return memo[amount] // Plus one as it's marked as a valid coin change
}

// min returns the minimum number between the two
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
