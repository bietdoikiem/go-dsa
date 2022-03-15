<p align="center"><img src="https://i.ibb.co/qxJQw1B/image.png" width="650" /></p>

# No. 70 - Climbing Stairs

## Core Concept

The problem indicates only 1 or 2 steps at a time is allow per climb until the
end of the stairs. For that reason, recursion can be used to recursively try
different steps and mark it as success if the step is successfuly reach the end
but not overstep it.

<p align="center"><img src="https://i.ibb.co/DMhVmSf/image.png" width="350" /></p>

This is one of the Dynamic Programming problem in which the recursion step must
be optimized, otherwise the solution will be flagged as Time Limit Exceeded
(TLE). For that reason, Top-Down approach (Memoization) will be used to
eliminate the overhead of recursion in calculated parameters of the recursive
function. To be specific, a Map will be used to store calculated result from
sub-recursion problem, and its result will be reused in other ones.

## Pseudocode

```text
FUNCTION climbStairs(n):
  memo := map{}
  ans := recursiveClimbStairs(0, n, memo)
  RETURN ans

FUNCTION recursiveClimbStairs(steps, n, memo):
  // Check if the result is already calculated/memoized
  IF val := memo[n]:
    RETURN val
  // Check if this is the end of the stair's steps
  IF n == 0:
    RETURN 1
  // Check for overstep
  IF n < 0:
    RETURN 0
  // Memoized calculated result of current recursion
  memo[n] = recursiveClimbStairs(1, steps-1) + recursiveClimbStairs(2, steps-2)
  RETURN memo[n]
```
