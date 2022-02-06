<p align="center"><img width="350" src="https://i.ibb.co/dg6ChNV/image.png" /></p>

# No. 53 - Maximum Subarray

## Core Concept

This problem can be solved effectively in linear time using only one pointer. We're going to use a kind of Greedy Algorithm to optimize for every value in each iteration; thus, the maximal value can be obtained at the end.

To be specific, we're going to calculate the sum of the array following the
iteration through each element by adding each element to the sum variable.
However, it would not be effective to sum the elements which decreases the value
of sum. For that reason, we're going to get rid of those elements if the current
iterated element has a higher value than the previously calculated sum by
setting the sum variable to the value of current element. In addition,
previously calculated sum with positive value (equals or greater than 0) will not be getting rid of due to
its extra value added.

<p align="center"><img width="350"
src="https://i.ibb.co/7n51wSd/Screenshot-20220206-224512-Samsung-Notes.jpg" /></p>

## Pseudocode

Time complexity: O(n)

Space complexity: O(1)

```
FUNCTION maxSubArray(nums []int):
  sum := nums[0]
  max := sum
  FOR i := 1 TO length(nums) - 1 DO:
    IF nums[i] > sum && sum < 0:
      sum = nums[i]
    ELSE:
      sum += nums[i]
    // Check maximum
    IF sum > max:
      max = sum
  ENDFOR
  RETURN max
```
