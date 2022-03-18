package longestvalidparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".

+Example 2:
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".

+Example 3:
Input: s = ""
Output: 0

+Example 4:
Input: s = "(()(()"
Output: 2

+Example 5:
Input: "()(()"
Output: 2
*/

func Test_longestValidParentheses_1(t *testing.T) {
	assert := assert.New(t)
	s := "(()"
	ans := longestValidParentheses(s)
	assert.Equal(2, ans)
}

func Test_longestValidParentheses_2(t *testing.T) {
	assert := assert.New(t)
	s := ")()())"
	ans := longestValidParentheses(s)
	assert.Equal(4, ans)
}

func Test_longestValidParentheses_3(t *testing.T) {
	assert := assert.New(t)
	s := ""
	ans := longestValidParentheses(s)
	assert.Equal(0, ans)
}

func Test_longestValidParentheses_4(t *testing.T) {
	assert := assert.New(t)
	s := "(()(()"
	ans := longestValidParentheses(s)
	assert.Equal(2, ans)
}

func Test_longestValidParentheses_5(t *testing.T) {
	assert := assert.New(t)
	s := "()(()"
	ans := longestValidParentheses(s)
	assert.Equal(2, ans)
}
