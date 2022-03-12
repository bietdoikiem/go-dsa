<p align="center"><img src="https://i.ibb.co/ftRzQ9Z/image.png" width="350" /></p>

# No. 15 - 3Sum

## Core Concept

This problem is really straightforward if "Two Sum II" has been solved in prior
as it's pretty similar regarding the solution's concept.

First of all, let's just simply think of the problem the same way as the Two Sum
problem variants. To be specific, the sum of three numbers to be zero is actually
the sum of the other two numbers which will be the opposite of the first one to
make it zero. Thus, a hash map, two converging pointers technique can also be
used to solve this problem as it's now pretty similar to Two Sum problems.

For this problem, two converging pointers technique will be used to find the sum for the
opposite of the target, which is the most commonly used method for "Two Sum II"
problem. In addition, the technique is quite optimal due to its O(n) complexity
for finding the combination of two in sorted array.

Firstly, the array is required to be sorted (for the two pointers technique)
which costs O(nlogn) for the fastest sort algorithm in this case is Quick Sort.
Sequentially, an outer loop will have the responsibility to keep track of the
target (specifically ith index element starting from 0 in the array), and an
inner nested loop will manage two of pointers in order to make the sum of the
overall three elements equal to zero. In total, a complexity of O(N<sup>2</sup>)
is witnessed in the solution.

<p align="center"><img src="https://i.ibb.co/YkGQFBs/IMG-0050.jpg" width="350" /></p>

Having applied the technique to acquire the triplets, there are some edge cases
happened due to the constraints of the problem. In details, the triplets must be
completely unique from each other. As for that, if the two pointers technique is
only applied without taking care of the edge cases, duplicated is absolutely
found. For those reasons, there are THREE special checks that should be
integrated into the solution to avoid the duplication issue in triplets.

3Sum with 3Checks:

1. Still continue to CHECK for equivalent combinations whose sum is equal to
   target even after we've found one combination in the iteration.
2. CHECK for duplicated target (neighbor) as it can produce duplicated
   triplets (Skip the duplicated target immediately)
3. CHECK for duplicated left, right pointers as it can also produce duplicated
   triplets (Skip and shift the pointers to their next ones)

## Pseudocode

```text
FUNCTION threeSum(nums):
  n := length(nums)
  var ans [][]int
  // Skip checking for less than three elements in an array
  IF n < 3:
    RETURN ans
  var l int
  var r int
  FOR i = 0 TO n - 1 DO:
    // Skip previously-known element
    IF i > 0 && nums[i-1] == i:
      CONTINUE
    // Define pointers
    l = i + 1
    r = n - 1
    FOR l < r DO:
      // Skip checking duplicated left/right pointers
      IF l > i+1:
        IF nums[l] == nums[l-1]:
          l++
          CONTINUE
      IF r < n-1:
        IF nums[r] == nums[r+1]:
          r--
          CONTINUE
      // Two converging pointers technique
      sum := nums[i] + nums[l] + nums[r]
      IF sum == 0:
        append(ans, [nums[i], nums[l], nums[r])
        l++
        r--
      ELSE IF sum > 0:
        r--
      ELSE:
        l++
    ENDFOR
  ENDFOR
  RETURN ans
```
