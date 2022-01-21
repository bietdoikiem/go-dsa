package lists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_linked_list(t *testing.T) {
	// Init Assertion lib
	assert := assert.New(t)
	list := NewLinkedList[int]()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	fmt.Print("List items after insertion: ")
	displayIntList(list)
	assert.Equal(list.size, 5, "Size should be 5 after insertion.")
	// DeleteByKey Test
	// 1st Key deleted
	d1, err := list.DeleteByKey(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deleted value is %v \n", *d1)
	// 4th key deleted
	d2, err := list.DeleteByKey(4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deleted value is %v \n", *d2)
	_, err = list.DeleteByKey(4) // Delete un-existing value
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("List items after delete by keys 1 and 4: ")
	displayIntList(list)
	assert.Equal(list.size, 3, "Size should be 3 after removing two keys.")
	// DeleteByIndex Test
	// Index 1st text
	d4, err := list.DeleteByIndex(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deleted value is %v \n", *d4)
	fmt.Print("List items after delete by index 0th: ")
	displayIntList(list)
	assert.Equal(list.size, 2, "Size should be 2 after removing the 0th index element")
}

// displayIntList shows a list of all integer items of the given list
func displayIntList(l LinkedList[int]) {
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
