<p align="center"><img width="350" src="https://i.ibb.co/ByS5Y3T/image.png" /></p>

# No. 238 - Product of Array Except Self

## Core Concept

The problem can be really easy if it had not been restricted from using division
operator. For that reason, we're going to solve the problem using
only-multiplication method!

Let's think that the except-self product of the nums[i] can be calculated using
the product of all the elements which sits behind nums[i] and the one after
nums[i].

<p align="center"><img width="350" src="https://i.ibb.co/cvY44gq/Screenshot-20220205-234436-Samsung-Notes.jpg"/></p>

As indicated by the picture above, the product of all elements behind nums[i] is
called Prefix, and its counterpart is call Postfix. We're going to calculate
the Prefix array, and the Postfix array as the product of i index at those two
arrays equal the except-self product of num[i].

To calculate the Prefix array, we see that the product of elements behind _ith_
index is actually the product of element _(i-1)th_ with the prior
result of \*(i-1)th index at the Prefix array.

<p align="center"><img width="350" src="https://i.ibb.co/Jp3j0wC/Screenshot-20220206-001521-Samsung-Notes.jpg"/></p>

To calculate the Postfix array, the same process goes with the Prefix one except
we're going to do it in a reverse way by starting from the last index.

<p align="center"><img width="350" src="https://i.ibb.co/M57NXyJ/Screenshot-20220206-001706-Samsung-Notes.jpg"/></p>

**Note**:

> The 0th index of the Prefix array and the (n-1)th index of the Postfix array is initialized as 1 since there are no prior results for calculation.

Finally, we got the output/answer array by multiplied the result of n elements
in Prefix and Postfix array.

<p align="center"><img width="350" src="https://i.ibb.co/w053FjM/Screenshot-20220206-001859-Samsung-Notes.jpg"/></p>

## Pseudocode

Time complexity: O(2n) => O(n)

Space complexity: O(2n) => O(2n) (excluding output array)

```
FUNCTION productExceptSelf(nums):
  n := length(nums)
  // Initialize length of n arrays
  outputArr := [0..0]
  prefixArr := [0..0]
  postfixArr := [0..0]
  // Initialize pre/post arrays for first/last element of those arrays
  prefixArr[0] = 1
  postfixArr[0] = 1
  // Calculate prefix array
  FOR i := 1 TO n - 1 DO:
    prefixArr[i] = nums[i-1] * prefixArr[i-1]
  ENDFOR
  // Pre-compute first/last element of output array
  outputArr[n-1] = prefixArr[n-1] * postfixArr[n-1]
  FOR i := n - 2 TO 0 DO:
    postfixArr[i] = nums[i+1] * postfixArr[i+1]
    outputArr[i] = prefixArr[i] * postfixArr[i]
  ENDFOR
  RETURN outputArr
```
