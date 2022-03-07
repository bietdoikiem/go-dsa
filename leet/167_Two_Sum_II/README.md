<p align="center"><img src="https://i.ytimg.com/vi/cQ1Oz4ckceM/maxresdefault.jpg" width="350" /></p>

# No. 167 - Two Sum II

## Core Concept

This problem is extremely straighforward if one may have encoutered converging
pointers solution before.

Specifically, there will two pointers: one at the left-most, and one at the
right most. Since the sorted array is given, it is assumed that the left-most
element has the lowest value, which is oppesite the right-most. Thus, after
calculating and comparing the sum of two pointer elements, if the value is
larger than target, the right-most index will be shifted to left (by minus one).
On the other hand, if the it is smllaer than the target, the left-most will be
shifted to the left (by plus one).

<p align="center"><img src="https://i.ibb.co/VM1P31D/image.png" width="350" /></p>

## Pseudocode

```text
FUNCTION twoSum(nums []int, target int):
  n := length(nums)
  l := 0 // Left-most
  r := n - 1 // Right-most
  WHILE l < r DO:
    sum := nums[l] + nums[r]
    IF sum == target:
      RETURN [ + 1, r + 1]
    IF sum > target:
      r--
    ELSE:
      l++
  ENDWHILE
```
