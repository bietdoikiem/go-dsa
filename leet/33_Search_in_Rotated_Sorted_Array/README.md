<p align="center"><img
src="https://res.cloudinary.com/practicaldev/image/fetch/s--woyQQuRh--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/7nfo9kh59xxybpwpz45q.png"
width="350" /></p>

# No. 33 - Search in Rotated Sorted Array

## Core Concept

This questions requires us to handle the edge cases really specifically in order
to achieve O(logn) time complexity algorithm.

Upon approaching the problem, the O(logn) requirement for algorithms has already
given us a hint for the solution. For that reason, it's best to think about
Binary Search for this kind of Rotated sorted array.

Firstly, we need to check for 2 specific cases that can occur when choosing a
pivot in the rotated sorted array.

<p align="center"><img src="https://i.ibb.co/PhLB0nv/image.png" width="350"></p>

According to the above picture, we can see that the middle pivot is bigger
than the left-most element, larger than the target, but the target is obviously
smaller than the left-most element. Therefore, the previous elements up to the
pivot has already been rotated. For that reason, it is necessary to shift
the left-most index to the right of the current middle pivot (l = m + 1) in order to correctly
narrow the search range.

<p align="center"><img src="https://i.ibb.co/rG3nwDH/image.png" width="350"></p>

On the second case, indicating in the above image, it is completely opposite the
first case in which the left-most element, and the target is larger than the
middle pivot, and the target is also larger than the left-most element. That
case happens when the array is not rotated up to the middle pivot.

Moreover, to avoid skipping the elements while choosing left-most, right-most
and middle pivot elements, there will be in total three conditions to check for
if those elements are equal to the target at the beginning of the iteration.

## Pseudocode

```text
FUNC search(nums, target):
  n := length(nums)
  // Skip-checkin for one-element array
  IF n == 1:
    IF nums[0] == target:
      RETURN nums[0]
    RETURN -1
  // Pointers declaration
  l = 0
  r = n - 1
  FOR l < r:
    // Check current indexed elements
    IF nums[m] == target:
      RETURN m
    IF nums[l] == target:
      RETURN l
    IF nums[r] == target:
      RETURN m

    // Check for case, where pivot is greater than left-most and target but target is
    smaller than left-most
    IF nums[m] > nums[l] && nums[m] > target && target < nums[l]:
      // Shift left pointer
      l = m + 1
      CONTINUE
    // Check for case, where pivot is smaller than left-most and target; target
    // is greater than left-most
    IF nums[m] < nums[l] && nums[m] < target && target > nums[l]:
      r = m - 1
      CONTINUE

    // Else, do the binary search as usual since the current array is sorted
    IF num[m] > target:
      r = m - 1
    ELSE:
      l = m + 1
  ENDFOR
  // If not found
  RETURN -1
```
