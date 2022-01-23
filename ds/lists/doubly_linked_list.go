package lists

type DoublyLinkedListNode[T comparable] struct {
	value T
	next  *T
	prev  *T
}

// NewDoublyLinkedListNode instantiates an instance of doubly linked list node
func NewDoublyLinkedListNode[T comparable](value T) DoublyLinkedListNode[T] {
	return DoublyLinkedListNode[T]{value: value}
}

type DoublyLinkedList[T comparable] struct {
	size int
	head *T
	tail *T
}

// NewDoublyLinkedList instantiates an empty doubly linked list
func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}
