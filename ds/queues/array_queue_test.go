package queues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_array_queue(t *testing.T) {
	assert := assert.New(t)
	queue := NewArrayQueue[int](5)
	// Enqueue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	err := queue.Enqueue(6)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(5, queue.size, "Queue's size should be 5 after enqueuing the listed items.")
	// Dequeue
	first := queue.items[0]
	d1, err := queue.Dequeue()
	assert.Equal(4, queue.size, "Queue's size should be 4 after dequeuing.")
	assert.Equal(first, d1, "Dequeued item should be the same as the first item.")
	fmt.Print("Current queue: ")
	displayArrayQueue(queue)
}

func displayArrayQueue[T any](q ArrayQueue[T]) {
	for _, v := range q.items {
		fmt.Print(v, "")
	}
	fmt.Println("")
}
