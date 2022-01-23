package lists

import (
	"errors"
	"fmt"
)

// a SinglyLinkedListNode contains the actual value and the reference to the next node
type SinglyLinkedListNode[T comparable] struct {
	value T
	next  *SinglyLinkedListNode[T]
}

// NewSinglyLinkedListNode creates a new node
func NewSinglyLinkedListNode[T comparable](value T) SinglyLinkedListNode[T] {
	return SinglyLinkedListNode[T]{value: value}
}

// a SinglyLinkedList contains a list of nodes and a reference to the head
type SinglyLinkedList[T comparable] struct {
	size int
	head *SinglyLinkedListNode[T]
}

// NewSinglyLinkedList initiates an empty linked list
func NewSinglyLinkedList[T comparable]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{}
}

// Insert inserts a new node to the linked list
// Time complexity:
// 1. O(1) - Insert at root
// 2. O(n) - Insert at middle
// 3. O(n) - Insert at tail (can be improved to O(1) if tail pointer is tracked)
func (l *SinglyLinkedList[T]) Insert(value T) {
	inserted := NewSinglyLinkedListNode(value)
	// 1. Insert at head
	if l.head == nil {
		l.head = &inserted
		l.size += 1
		return
	}
	// 2. Insert at the middle or end
	var prev *SinglyLinkedListNode[T]
	cur := l.head
	// Traverse until the current is nil
	// Then, insert a new node at that position
	for cur != nil {
		prev = cur
		cur = cur.next
	}
	prev.next = &inserted
	l.size += 1
}

// DeleteByKey deletes a node in the list by its value
// Time complexity:
// 1. O(1) - Delete at first key
// 2. O(n) - Delete at middle
// 3. O(n) - Delete at tail
func (l *SinglyLinkedList[T]) DeleteByKey(value T) (*T, error) {
	// Check for empty list
	if l.head == nil {
		return nil, errors.New("Linked list is empty")
	}
	var deleted *T
	// Delete at head-only list
	if l.head.next == nil {
		deleted = &l.head.value
		l.head = nil // Set to nil for garbage collector
		l.size -= 1
		return deleted, nil
	}
	var prev *SinglyLinkedListNode[T]
	cur := l.head
	// Delete at head
	if cur.value == value {
		deleted = &cur.value
		l.head = cur.next
		cur = nil // Set to nil for garbage collector
		l.size -= 1
		return deleted, nil
	}
	// Delete at body
	for cur != nil {
		if cur.value == value {
			deleted = &cur.value
			prev.next = cur.next
			cur = nil // Set to nil for garbage collector
			l.size -= 1
			return deleted, nil
		}
		prev = cur
		cur = cur.next
	}
	return nil, fmt.Errorf("Cannot delete! The key %v does not exist", value)
}

// DeleteByIndex deletes a node in the list by its index/position
// Time complexity:
// 1. O(1) - Delete at index 0th
// 2. O(n) - Delete at middle index
// 3. O(n) - Delete at last index (tail)
func (l *SinglyLinkedList[T]) DeleteByIndex(index int) (*T, error) {
	// Check for empty list
	if l.head == nil {
		return nil, errors.New("Linked list is empty")
	}
	// Check out-of-bound index
	if index >= l.size {
		return nil, errors.New("Index is out of bounds")
	}
	var deleted *T
	// Delete at head-only
	if index == 0 && l.head.next == nil {
		deleted = &l.head.value
		l.head = nil
		l.size -= 1
		return deleted, nil
	}
	var prev *SinglyLinkedListNode[T]
	cur := l.head
	// Delete at head
	if index == 0 {
		deleted = &cur.value
		l.head = cur.next
		cur = nil
		l.size -= 1
		return deleted, nil
	}
	// Delete at body
	i := 1
	// Update pointers to index 1st element
	prev = cur
	cur = cur.next
	for i < l.size {
		if i == index {
			deleted = &cur.value
			prev.next = cur.next
			cur = nil
			l.size -= 1
			break
		}
		prev = cur
		cur = cur.next
	}
	return deleted, nil
}
