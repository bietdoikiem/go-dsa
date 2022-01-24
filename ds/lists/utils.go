package lists

import "fmt"

// DisplaySinglyLinkedList shows a list of all integer items of the given list
func DisplaySinglyLinkedList[T comparable](l SinglyLinkedList[T]) {
	if l.head == nil {
		fmt.Print("List is currently empty!")
		return
	}
	cur := l.head
	for cur != nil {
		fmt.Print(cur.value, " ")
		cur = cur.next
	}
	fmt.Println("")
}

// DisplaySinglyLinkedList shows a list of all integer items of the given list
func DisplayDoublyLinkedList[T comparable](l DoublyLinkedList[T]) {
	if l.head == nil {
		fmt.Print("List is currently empty!")
		return
	}
	cur := l.head
	for cur != nil {
		fmt.Print(cur.value, " ")
		cur = cur.next
	}
	fmt.Println("")
}
