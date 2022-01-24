package stacks

import "errors"

type ArrayStack[T any] struct {
	capacity uint
	size     uint
	items    []T
}

// NewArrayStack function returns a pointer to a newly created Stack using array
func NewArrayStack[T any](capacity uint) ArrayStack[T] {
	return ArrayStack[T]{capacity: capacity, size: 0, items: make([]T, capacity)}
}

// Push pushes the specified item onto the top of the stack
// Time complexity (amortized): O(1) -> As the slice does not grow out of capacity
func (s *ArrayStack[T]) Push(item T) error {
	if s.IsFull() {
		return errors.New("stack error: number of items has reached capacity")
	}
	s.items = append(s.items, item)
	s.size += 1 // Update current size
	return nil
}

// Pop removes the top item in the stack
// Time complexity: O(1) -> Re-slicing only
func (s *ArrayStack[T]) Pop() (T, error) {
	var poppedItem T
	if s.IsEmpty() {
		return poppedItem, errors.New("stack error: the stack is empty")
	}
	poppedItem = s.items[s.size-1]
	s.items = s.items[:s.size-1] // Remove last element O(1)
	s.size -= 1
	return poppedItem, nil
}

// Peek reviews the last added item in the stack
// Time complexity: O(1)
func (s *ArrayStack[T]) Peek() T {
	return s.items[s.size-1]
}

// IsEmpty checks if the stack contains no elements
func (s ArrayStack[T]) IsEmpty() bool {
	return s.size == 0
}

// IsFull checks if the stack is full
func (s ArrayStack[T]) IsFull() bool {
	return s.size == s.capacity
}
