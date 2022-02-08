package reverselinkedlist

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse a Linked List
func reverseList(head *ListNode) *ListNode {
	// Check empty list
	if head == nil {
		return head
	}
	// Check only-1 list
	if head.Next == nil {
		return head
	}
	// Define 3 pointers
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev // Reverse link
		prev = cur
		cur = next
	}
	cur = prev
	return cur
}
