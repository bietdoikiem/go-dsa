<p align="center"><img width="350" src="https://i.ibb.co/bmxHdhr/image.png" /></p>

# No. 152 - Maximum Product Subarray

## Core Concept

In this problem, we're going to solve the problem in 2 practical ways in order
to compare and contrast their performances.

### 1. Exhaustive Search/Brute Force approach

In this approach, we're going to iterate through all possible combinations of
subarray using an addition nested loop.

<p align="center"><img width="350" src="https://i.ibb.co/jvGFm2T/Screenshot-20220207-110153-Samsung-Notes.jpg" /></p>

Upon iteration through combinations, we'll multiply those elements altogether to get the product and keep track of the
maximum value based on the product result!

### 2. Optimal approach

Since we're calculating the product of sub-array, we can see that the result of
each product can be changed rapidly from largest to the smallest due to
negativity or '0' value. Moreover, smallest negative value of a subarray can turn out to
be the largest if the upcoming product with the next negative element.

For that reason, we'll want to keep track of both the current minimum and the maximum of
previously-calculated product in order to maximize the product in different
cases.

Nevertheless, the 0-value element may ruin the products of the current
min and max. To deal with the given edge case, since we're already obtain the
maximum value of any single elements in advanced, we'll confidently set the 0
value to be 1 (neutral value for product) to not disrupt the current product and
reset/skip the harming '0' value.

<p align="center"><img width="350" src="https://i.ibb.co/3rmVL5g/Screenshot-20220208-095015-Samsung-Notes.jpg" /></p>

## Pseudocode

### 1. Exhaustive Search/Brute Force approach

Time complexity: O(n<sup>2</sup>)

Space complexity: O(1)

```text
FUNCTION maxProduct(nums):
  n := length(nums)
  max := 0
  prod := 0
  FOR i := 0 TO n - 1 DO:
    prod = nums[i]
    // Check max of single-element product
    IF prod > max:
      max = prod
    FOR j := i + 1 TO n - 1 DO:
      prod *= nums[j]
      IF prod > max:
        max = prod
    ENDFOR
  ENDFOR
  RETURN max
```

### 2. Dynamic Programming optimal approach

Time Complexity: O(n)

Space Complexity: O(1)

```text
FUNCTION maxProduct(nums):
  n := length(nums)
  // Handle case: only-1 element
  IF n == 1:
    RETURN nums[0]
  // Find the maximum-value element in the array
  max := nums[0]
  FOR i := 1 TO n - 1:
    IF nums[i] > max:
      max = nums[i]
  // Store current max and min of subarray product (Dynamic Programming)
  curMax := nums[0]
  curMin := nums[0]
  curMaxProd := 1 // Product of current maximum and current
  curMinProd := 1 // Product of current minimum and current
  FOR i := 1 TO n - 1 DO:
    // Handle edge case '0'
    IF nums[i] == 0:
      // Reset current max and min to 1 to avoid 0 value
      curMax = 1
      curMin = 1
      CONTINUE
    // Calculate to choose current max and min between 3 selection
    curMaxProd = curMax * nums[i]
    curMinProd = curMin * nums[i]
    curMax = max(max(curMaxProd, curMinProd), nums[i])
    curMin = min(min(curMaxProd, curMinProd), nums[i])
    IF curMax > max {
      max = curMax
    }
  ENDFOR
  RETURN max
```
