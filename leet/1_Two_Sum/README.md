<p align="center"><img width="350" src="https://i.ibb.co/166j0fh/image.png" /></p>

# No. 1 - Two Sum

## Core Concept

This problem really straightforward if do the O(n^2) algorithm to find
appropriate pair by calculating the difference from the current element.
However, that would be a waste of time to just search only 1 element in linear
time. For that reason, we'll be using Map data structure to enhance the search
time to constant O(1) which will aid the process of finding appropriate pair to
O(n) only.

## Pseudocode

Time complexity: O(n)

Space Complexity: O(n) (auxiliary map)

```
FUNCTION twoSum(nums, target):
  map := DECLARE map whose key is integer and value is integer
  n := length(nums)
  FOR i := 0 TO n - 1 DO:
    diff := target-nums[i]
    IF map[diff] exists:
      RETURN [i, map[diff]]
    // Else, keep track of current element's index in map
    records[nums[i]] = i
  RETURN empty array
```
