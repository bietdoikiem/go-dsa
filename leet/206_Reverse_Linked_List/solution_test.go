package reverselinkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+ Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

+ Example 2:
Input: head = [1,2]
Output: [2,1]

+ Example 3:
Input: head = []
Output: []
*/

func Test_reverseList_1(t *testing.T) {
	assert := assert.New(t)
	// Declare linked list
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printList(head)
	newHead := reverseList(head)
	fmt.Println("After reverse: ")
	printList(newHead)
	// Assertions
	assert.Equal(5, newHead.Val, "New head's value should be 5 after reverse.")
}

func Test_reverseList_2(t *testing.T) {
	assert := assert.New(t)
	// Declare linked list
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	printList(head)
	newHead := reverseList(head)
	fmt.Println("After reverse: ")
	printList(newHead)
	// Assertions
	assert.Equal(2, newHead.Val, "New head's value should be 2 after reversing.")
}

func Test_reverseList_3(t *testing.T) {
	assert := assert.New(t)
	// Declare linked list
	var head *ListNode
	printList(head)
	newHead := reverseList(head)
	fmt.Println("After reverse: ")
	printList(newHead)
	// Assertions
	assert.Equal(head, newHead, "New head's value should be nil.")
}

// printList prints the list of node's values
func printList(head *ListNode) {
	if head == nil {
		fmt.Print("List is currently empty!")
		return
	}
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println("")
}
