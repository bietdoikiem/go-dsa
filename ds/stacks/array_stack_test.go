package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_array_stack(t *testing.T) {
	assert := assert.New(t) // Init assertion library
	stack := NewArrayStack[int](5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	// Assert Push
	assert.Equal(stack.size, 5, "Stack size should be 5.")
	assert.Equal(stack.top, 4, "Top item index should be 4.")
	// Assert Peek
	assert.Equal(stack.Peek(), stack.items[stack.top], "Peeked item should be the top one in the array.")
	// Assert Pop
	topItem := stack.items[stack.top]
	popItem := stack.Pop()
	assert.Equal(stack.size, 4, "Stack size should be 4 after popping last item.")
	assert.Equal(stack.top, 3, "Top item index should be 4 after popping last item.")
	assert.Equal(popItem, topItem, "Popped item should be equal to the previously top one.")
}
