package lists

import "errors"

type DoublyLinkedListNode[T comparable] struct {
	Node[T]
	next *DoublyLinkedListNode[T]
	prev *DoublyLinkedListNode[T]
}

// NewDoublyLinkedListNode instantiates an instance of doubly linked list node
func NewDoublyLinkedListNode[T comparable](value T) DoublyLinkedListNode[T] {
	return DoublyLinkedListNode[T]{Node: Node[T]{value: value}}
}

type DoublyLinkedList[T comparable] struct {
	size int
	head *DoublyLinkedListNode[T]
	tail *DoublyLinkedListNode[T]
}

// NewDoublyLinkedList instantiates an empty doubly linked list
func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}

// Insert inserts a new node to the linked list
// Time complexity: O(1)
func (l *DoublyLinkedList[T]) Insert(value T) {
	// Init new node
	inserted := NewDoublyLinkedListNode(value)
	// Insert at head
	if l.size == 0 && l.head == nil {
		l.head = &inserted
		l.tail = &inserted
		l.size += 1
		return
	}
	// Insert at tail
	cur := l.tail
	cur.next = &inserted // Update next
	inserted.prev = cur  // Update prev
	l.tail = &inserted   // Update tail
	l.size += 1
}

// DeleteByKey deletes a node from the linked list by its key
func (l *DoublyLinkedList[T]) DeleteByKey(value T) (*T, error) {
	// Check empty list
	if l.IsEmpty() {
		return nil, errors.New("doublylinkedlist error: list is currently empty")
	}
	var deleted *T // Init deleted pointer
	// Delete at root-only
	if l.size == 1 && l.head.value == value && l.head.next == nil {
		deleted = &l.head.value
		l.head = nil
		l.tail = nil
		l.size -= 1
		return deleted, nil
	}
	// Delete at head
	if value == l.head.value {
		deleted = &l.head.value
		// Update pointer of the next element
		l.head = l.head.next
		l.head.prev = nil
		return deleted, nil
	}
	// Delete at body (start after head)
	prev := l.head
	cur := l.head.next
	for cur != nil {
		if cur.value == value {
			deleted = &cur.value
			next := cur.next // Next of current pointer
			prev.next = next
			next.prev = prev
			return deleted, nil
		}
		prev = cur
		cur = cur.next
	}
	return nil, errors.New("doublylinkedlist error: key does not exist")
}

func (l *DoublyLinkedList[T]) DeleteByIndex(index int) {

}

// IsEmpty checks if the linked list is currently empty
func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.size == 0 && l.head == nil && l.tail == nil
}
