package queues

import "errors"

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
func (q *ArrayQueue[T]) Enqueue(item T) error {
	if q.isFull() {
		return errors.New("Cannot enqueue. The queue is full")
	}
	q.items[q.size] = item
	q.size += 1
	return nil
}

func (q *ArrayQueue[T]) Dequeue() error {
	if q.isEmpty() {
		return errors.New("Cannot dequeue. The queue is empty")
	}
	// Shift all items to the left
	for i := 1; i < q.size; i++ {
		q.items[i-1] = q.items[i]
	}
	// Remove last element
	return nil
}

// isEmpty checks if the queue is currently empty
func (q ArrayQueue[T]) isEmpty() bool {
	return q.size == 0
}

// isFull checks if the queue has already been full
func (q ArrayQueue[T]) isFull() bool {
	return q.size == q.capacity
}
