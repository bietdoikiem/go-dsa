package besttimetobuyandsellstocks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

+Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/

func Test_maxProfit_1(t *testing.T) {
	assert := assert.New(t)
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println("Result is: ", result)
	assert.Equal(5, result, "Maximized profit should be 5.")
}

func Test_maxProfit_2(t *testing.T) {
	assert := assert.New(t)
	prices := []int{7, 6, 4, 3, 1}
	result := maxProfit(prices)
	fmt.Println("Result is: ", result)
	assert.Equal(0, result, "Maximized profit should be 0.")
}
