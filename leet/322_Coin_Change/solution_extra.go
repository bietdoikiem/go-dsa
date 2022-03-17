package coinchange

// coinChangeBottomUp calculate the number of coins to use as change in Bottom-Up DP approach
func coinChangeBottomUp(coins []int, amount int) int {
	// Check if amount is zero
	if amount == 0 {
		return 0
	}
	bound := amount + 1 // Upper Boundary for current amount
	cache := make([]int, amount+1)
	// Initialize max value for all values of amount
	for a := 1; a <= amount; a++ {
		cache[a] = bound
	}
	// Pre-initialized for coin types that matches the amount
	for _, c := range coins {
		if c > amount {
			continue
		}
		cache[c] = 1
	}
	// NOTE: Bottom-up means that's we must start from the buttom, opposite to the memoization/top-down DP approach recursion
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			// Check if current coin does not fit the current amount
			if a-c <= 0 {
				continue
			}
			cache[a] = min(cache[a], 1+cache[a-c])
		}
	}
	// Retrieve answer if it's viable
	if cache[amount] < bound {
		return cache[amount]
	}
	return -1
}
