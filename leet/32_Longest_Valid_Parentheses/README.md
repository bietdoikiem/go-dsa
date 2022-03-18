<p align="center"><img
src="https://media.giphy.com/media/j7kSTil1J4gQQxOkXL/giphy.gif" width="450" /></p>

# No. 32 - Longest Valid Parentheses

## Core Concept

In this problem, a stack can be used as usual to deal with brackets/parentheses
problem due to their nature of presence - one opening and one closing.

The stack will be used to match a valid pair of bracket or nested pair due to
its FIFO nature. Since only the length of the substring is required, a greedy
approach O(n) can be utilized to obtain the longest substring with a stack by
keep tracking of the open brackets inside it.

However, in this problem, a "()()" is even counted as two valid
pairs which results in the answer of 4 valid brackets. For that reason, the
second to last matched bracket index should be used to calculated the length of
substring in order not to miss the first valid bracket.

One more problem to think about is that, if the stack has been popped already
with no opening bracket left (in case "()()"), then how can the result of the
second to last one can be used? The answer is that an extra element "-1" can be
inserted at the beginning of the stack to handle this case.

Nevertheless, there's another problem encountered when using an extra element is
that what if there's an closing bracket when the stack is only left with the
extra element. In that case, the value of the extra element can be reset by
updating its value with the current traversed index.

## Pseudocode

```text
FUNC longestValidParentheses(s):
  n := length(s)
  // Check for empty string
  IF s == "":
    RETURN ""
  var longest int
  stack := [-1] // Init -1 outer value
  FOR i := 0 TO n DO:
    IF s[i] == '(':
      stack = append(stack, i)
    ELSE:
      // Check if there's only one outer value in stack
      IF length(stack) == 1:
        stack[0] = i // Reset outer value
        CONTINUE
      // Else pop stack && update longest string
      stack = stack[:length(stack)-1]
      longest = max(longest, i-stack[length(stack)-1])
  ENDFOR
  RETURN longest
```
