<p align="center"><img
src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg"
width="400" /></p>

# No. 11 - Container With Most Water

## Core Concept

The problem specifies that max area to contain the water must be calculated
without slanting the container (ie. disallow product of two different heights,
only multiply with the minimum between the two).

Given the problem, the formula to calculate max area can easily be formed:
**(i<sup>th></sup> - j<sup>th</sup>)\*min(h1, h2)**, in which i is the index of the first bar, and
j is the second one; h1 and h2 is the height of two different bars.

**Note**:

> For this problem, Brute Force method, obviously O(n<sup>2</sup>)
> complexity, is unfortunately not allowed as it's
> restricted with Time Limit Exceeded (TLE).

Since the distance between the two vertical bars/lines is as important as the
height of those two, it has given out a hint that the O(n) algorithm for is
this problem should be using a sliding-window technique, in which left and right
pointers are initialized at two ends of the array. One of the pointer is shifted
towards the middle if it's smaller than its counterpart.

<p align="center"><img src="https://i.ibb.co/wS0w2nc/image.png" width="350" /></p>

## Pseudocode

Time Complexity: O(n)

Space Complexity: O(1)

```text
FUNCTION maxArea(height):
  n := length(height)
  var maxArea int
  var curArea int
  l := 0
  r := n - 1
  WHILE l < r DO:
    curArea = (j-i)*min(height[i], height[j])
    maxArea = max(maxArea, curArea)
    IF height[i] > height[j]:
      r--
    ELSE:
      l++
  ENDWHILE
  RETURN maxArea
```
