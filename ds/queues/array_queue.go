package queues

import (
	"errors"
	"fmt"
)

type ArrayQueue[T any] struct {
	capacity int
	size     int
	items    []T
}

// NewArrayQueue instantiates a new queue implemented with Array
func NewArrayQueue[T any](capacity int) ArrayQueue[T] {
	return ArrayQueue[T]{capacity: capacity, size: 0, items: make([]T, capacity)}
}

// Enqueue adds a new item to the queue
// Time Complexity: O(1)
func (q *ArrayQueue[T]) Enqueue(item T) error {
	if q.isFull() {
		return fmt.Errorf("queue error: cannot enqueue %v as the queue is full", item)
	}
	q.items[q.size] = item
	q.size += 1
	return nil
}

// Dequeue dequeues the first item from the queue
// Time complexity: O(n)
func (q *ArrayQueue[T]) Dequeue() (T, error) {
	// Init empty struct
	var empty T
	// Check empty queue
	if q.isEmpty() {
		return empty, errors.New("queue error: the queue is empty")
	}
	// Retrieve 0th dequeued item
	dequeued := q.items[0]
	// Shift elements to the left
	for i := 1; i < q.size; i++ {
		q.items[i-1] = q.items[i]
	}
	q.items[q.size-1] = empty // Set empty struct for the last un-shifted items
	q.size -= 1
	return dequeued, nil
}

// Peek peeks the first item in the queue
// Time complexity: O(1)
func (q ArrayQueue[T]) Peek() (T, error) {
	var empty T
	if q.isEmpty() {
		return empty, errors.New("queue error: the queue is empty")
	}
	return q.items[0], nil
}

// isEmpty checks if the queue is currently empty
func (q ArrayQueue[T]) isEmpty() bool {
	return q.size == 0
}

// isFull checks if the queue has already been full
func (q ArrayQueue[T]) isFull() bool {
	return q.size == q.capacity
}
