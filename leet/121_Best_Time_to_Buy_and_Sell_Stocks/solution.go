package besttimetobuyandsellstocks

// maxProtfit returns the maximum profit during trading period
// Time complexity: O(n)
func maxProfit(prices []int) int {
	n := len(prices)
	buyIdx := 0
	sellIdx := 1
	var cashOut int
	max := 0 // Current max profit
	for buyIdx < n && sellIdx < n {
		if prices[sellIdx] < prices[buyIdx] {
			buyIdx = sellIdx // Buy at the dip!
		} else {
			cashOut = prices[sellIdx] - prices[buyIdx]
			if cashOut > max {
				max = cashOut
			}
		}
		sellIdx++
	}
	return max
}
