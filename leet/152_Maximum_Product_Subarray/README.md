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

## Pseudocode

### 1. Exhaustive Search/Brute Force approach

Time complexity: O(n<sup>2</sup>)

Space complexity: O(1)

```text
FUNCTION maxProduct(nums):
  n := len(nums)
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
