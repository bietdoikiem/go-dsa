<p align="center"><img src="https://c.tenor.com/wjS2sXen8iMAAAAM/stonks-up-stongs.gif"></p>

# Leetcode no. 121

## Core concept

In this problem, we are required to obtain the maximum profit out of a
arbitrary-size time window. As the time is linear (uni-directional) we can only
traverse, and check the buy & sell price at the front.

As for that reason, we'll use a 2-pointer technique (a variant of Sliding Window
Technique) to maximize out profit.

First of all, we'll use a pointer (called B pointer) to keep track of the price
index at which we have bought the stock. In sequence, another S pointer is used
to indicate the index of sell price. Upon moving forward the price array by
incrementing S, whenever we encounter a lower price than
the price at B index, we immediately update the pointer of the B index to be at
that price index (BUY THE DIP!). Then, we'll try to calculate the profit earned from
selling the stock and update the current maximum profit variable. In the end of
iterations, we'll return the maximum variable.

## Pseudocode

Time complexity: O(n)

Space complexity: O(1)

```
FUNCTION maxProfit(prices):
  n := length(prices)
  buyIdx := 0
  sellIdx := 1
  max := 0
  cashOut := 0
  FOR i := 0 TO n - 1:
    IF prices[sellIdx] < prices[buyIdx]:
      buyIdx = sellIdx
    ELSE:
      cashOut = prices[sellIdx] - prices[buyIdx]
      IF cashOut > max:
        max = cashOut
    sellIdx++
  RETURN max
```
