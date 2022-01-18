package data_structures

/*
	Implementation of the Stack data structure using Golang
	Methods:
	+ Push
	+ Pop
	+ Peek
	+ Clear
*/

type ArrayStack[T any] struct {
	capacity int
	size     int
	top      int
	items    []T
}

// NewArrayStack function returns a pointer to a newly created Stack using array
func NewArrayStack[T any](capacity int) *ArrayStack[T] {
	return &ArrayStack[T]{capacity: capacity, size: 0, top: -1, items: make([]T, capacity)}
}

// Push pushes the specified item onto the top of the stack
func (s *ArrayStack[T]) Push(item T) {
	s.top += 1
	s.items[s.top] = item
	s.size += 1
}

// Pop removes the top item in the stack
func (s *ArrayStack[T]) Pop() T {
	popItem := s.items[s.top]
	s.items = s.items[0:s.top] // Remove last element
	s.size -= 1
	s.top -= 1
	return popItem
}

// Peek reviews the last added item in the stack
func (s *ArrayStack[T]) Peek() T {
	return s.items[s.top]
}
