package lists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doubly_linked_list(t *testing.T) {
	assert := assert.New(t)
	// Init
	list := NewDoublyLinkedList[int]()
	// Insert
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	DisplayDoublyLinkedList(list)
	fmt.Println("Current Size of inserted list", list.size)
	assert.Equal(1, list.head.value, "Head value should be 1.")
	assert.Equal(5, list.tail.value, "Tail value should be 5.")
	// DeleteByKey
	list.DeleteByKey(1)
	list.DeleteByKey(4)
	DisplayDoublyLinkedList(list)
	assert.Equal(2, list.head.value, "Head value should be 2 after deletion.")
	// DeleteByKey error
	_, err := list.DeleteByKey(10)
	if err != nil {
		fmt.Println(err)
	}
	// DeleteByIndex
	list.DeleteByIndex(0)
	DisplayDoublyLinkedList(list)
	_, err = list.DeleteByIndex(1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = list.DeleteByIndex(0)
	// Check latest deletion
	DisplayDoublyLinkedList(list)
	assert.Equal(0, list.size, "Size should be 0.")
}
