package stacks

import (
	"fmt"
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
	assert.Equal(uint(5), stack.size, "Stack size should be 5.")
	// Assert Peek
	assert.Equal(stack.items[stack.size-1], stack.Peek(), "Peeked item should be the top one in the array.")
	// Assert Pop
	topItem := stack.items[stack.size-1]
	popItem, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(uint(4), stack.size, "Stack size should be 4 after popping last item.")
	assert.Equal(topItem, popItem, "Popped item should be equal to the previously top one.")
	stack.Push(6) // Push again
	fmt.Println("Last Pushed Item:", stack.items[stack.size-1])
}
