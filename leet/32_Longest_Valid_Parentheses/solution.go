package longestvalidparentheses

// longestValidParentheses returns the longest valid parentheses in the given string
func longestValidParentheses(s string) int {
	n := len(s)
	// Check for empty
	if s == "" {
		return 0
	}
	// Init stack
	stack := []int{-1}
	var longest int
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
		case ')':
			// Check if there's only element left at the stack (the default outer value)
			if len(stack) == 1 {
				stack[0] = i // Reset default value
				continue
			}
			// Pop the matched bracket from the stack
			stack = stack[:len(stack)-1]
			longest = max(longest, i-stack[len(stack)-1])
		}
	}
	return longest
}

// max returns the maximum number between 2 arguments
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
