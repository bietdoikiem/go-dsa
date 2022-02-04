<p align="center"><img src="https://i.ibb.co/yf4QMth/image.png" /></p>

# No. 217 - Contains Duplicate

## Core concept

Consider to use a Set data structure for this type of problem, or
feature-similar data structures such as HashMap, HashSet to eliminate duplicates
at linear time.

## Pseudocode

Time complexity: O(n)

Space complexity: O(n) - n elements in auxiliary Map data structure

```
FUNCTION containsDuplicate(nums):
  map := DECLARE a map whose key and value is integer
  FOR i := 0 TO length(nums) - 1:
    IF map[nums[i]] > 0:
      RETURN true
    maps[nums[i]] += 1
  ENDFOR
  RETURN false
```
