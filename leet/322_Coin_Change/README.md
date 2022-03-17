<p align="center"><img src="https://i.ibb.co/6JBhfxV/image.png" width="650" /></p>

# No. 322 - Coin Change

## Core Concept

### 1. DFS & Backtracking (Top-Down DP)

Given the problem's indication, it is known that different type of coins can be
used repeatedly and variedly until those coins fit the given amount of money.

For that reason, DFS can be used to exhaustively search for the correct
combinations of coins that make up for the given amount from the list of coins.
In sequence, having completed the process of finding number of coins needed for
the amount in different combinations, the minimum number will be chosen as the
solution. Thus, the final solution will be achieved by aggregating the
recursive sub-problem results.

Moreover, a slight enhancement can be put into the algorithm by caching the
calculated result with an auxiliary array from subproblem using Top-Down Dynamic Programming.

### 2. Bottom-up DP

It is straight forward to think about the Bottom-Up DP solution if the Top-Down
DP has been coded beforehand, as the result will be gradually calculated and aggregated
from the bottom (amount 0) to the top (amount n).

## Pseudocode

### 1. Top-down DP

```text
FUNCTION coinChange(coins, amount):
  // Initialize memo
  memo := []
  // Pre-filled the cache with different types of coin amount
  FOR coin in coins:
    memo[coins] = 1
  ans := helper(coins, amount, memo)
  IF ans == BOUNDARY_VALUE:
    RETURN -1
  RETURN ans

FUNCTION helper(coins, amount, memo):
  IF memo[amount] != 0:
    RETURN memo[amount]
  // Check if it's already fit the given amount
  IF amount == 0:
    RETURN 1
  // Check for overfit
  IF amount < 0:
    RETURN 0
  var amountLeft int
  minCoins := BOUNDARY_VALUE
  curCoins := 0
  // Iterate different given type of coins for generating combinations
  FOR coin in coins DO:
    amountLeft = amount - coin
    // Check for overfit coin
    IF amountLeft < 0:
      CONTINUE
    // Check for cached result
    IF memo[amountLeft] != 0:
      curCoins = memo[amountLeft]
    ELSE:
      memo[amountLeft] = helper(coins, amountLeft, memo)
      curCoins = amountLeft
    // Check if cached/computed result is an invalid combination
    IF curCoins == - 1:
      CONTINUE
    minCoins = min(minCoins, curCoins)
  ENDFOR
  // Check if minCoins is an invalid minimum
  IF minCoins == BOUNDARY_VALUE:
    return minCoins
  // Save the minimum result
  memo[amount] = minCoins
  return memo[amount]
```

### 2. Bottom-up DP

```text
FUNCTION coinChange(coins, amount):
  boundary := amount + 1 // Boundary is the exceeded limit of the given amount
  memo := []
  // Pre-fill the cache with boundary values (for min check)
  FOR a := 1; a <= amount; a++ DO:
    memo[a] = boundary
  ENDFOR
  // Pre-fill only-one answer for type of coins that fit the cached amount
  FOR coin in coins DO:
    memo[coin] = 1
  ENDFOR

  // Iterate and generate combinations
  FOR a := 1; a <= amount; a++ DO:
    FOR coin in coins:
      // Skip if overfit or for case with only-one coin
      IF a-c <= 0:
        CONTINUE
      memo[a] = min(memo[a], 1+memo[a-c])
    ENDFOR
  ENDFOR

  // Check for valid minimum
  IF memo[amount] < boundary:
    RETURN memo[amount]
  RETURN -1
```
