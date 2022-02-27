<p align="center"><img src="https://i.ibb.co/sJT5ZWm/image.png" width="350" /></p>

# No. 153 - Find Minimum in Sorted Array

## Core Concept

Upon thing about the problem, there's an interesting hint that Leetcode has
given which is the time complexity of the solution must be O(logn). That
constraint has actually enabled me to think of using some kind of binary
operations to perform for the algorithm.

Having a look at the rotated sorted array, it can clearly be seen that whenever
there are _n_ number of rotations (_n_ != length(arr)), there are always 2
different sorted parts existed the in the array.

<p align="center"><img width="350" src="https://i.ibb.co/pW8fbjY/IMG-0027.jpg" /></p>

For that reason, if we simply pick a random element (for example at the middle
index) as a pivot, it will separate the array into a sorted and an unsorted part
from that element as indicated in the figure below.

<p align="center"><img src="https://i.ibb.co/BP7Qm43/D0-F1-D12-F-D999-40-F0-BB01-EB9877-E0-F896.jpg" width="350"/></p>

If we take a close notice we can see that if there's a "break" point in which
the element of the unsorted separated array sits next to the sorted one, the
minium of those two will be the answer (the minimum element of the rotated
sorted array)!

<p align="center"><img
src="https://i.ibb.co/bmjNh9d/CEB39-A57-3583-4798-B6-BB-9958-DF62-A3-A8.jpg"
width="350" /></p>

Therefore, a variation of binary search would possibly a good method for solving
this problem. In details, the middle pivot will be picked, which is equal to the average of the left-most and
right-most indices. Sequentially, the array will be checked if it is already sorted by
comparing the pivot with the left-most and right-most elements (pivot element
should be greater than the left one, and smaller than the right one). Given that
the unsorted part of the array is detected upon comparing, the left-most or
right-index will be shifted to the current middle depending on which side is
detected. The remaining array will repeat the process until we reach the point such that the
difference between left and right index is only one (they're besides each other), then the smaller element
will be picked between the two indices as the solution (minimum is find on the
"break" point). Last but not least, there's a condition that should always be
keep tracked is that if the left and the right indices are already sorted
according to the given sorted rule, then the left-most element will simply be
the minimum.

## Pseudocode

```text
FUNCTION findMin(nums []int)
  n := length(nums)
  // Skip checkin for one-element array
  IF n == 1:
    RETURN nums[0]
  l := 0 // Left-most pointer
  r := n - 1 // Right-most pointer
  var m int // Declare middle pointer (pivot)
  WHILE l < r DO:
    m = (l+r)/2 // Calculate middle pointer by averaging two pointers
    // Check if left and right pointers are close to each other
    IF absDiff(l, r) == 1:
      RETURN min(nums[l], nums[r])
    // Check if array is already sorted by comparing left and right pointers
    with middle one
    IF nums[l] < nums[m] && nums[m] < nums[r]:
      RETURN nums[l] // Leftmost is the min in sorted array
    // Check if left-side is unsorted
    IF nums[l] > nums[m]:
      r = m
    // Check if right-side is unsorted
    IF nums[r] < nums[m]:
      l = m
  ENDWHILE
  RETURN INF_NEGATIVE_VALUE // Invalid input case

FUNCTION absDiff(x int, y int):
  IF x > y:
    RETURN x - y
  RETURN y - x

FUNCTION min(x int, y int):
  IF x > y:
    RETUNR y
  RETURN X
```
