<p align="center"><img width="350" src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg" /></p>

# No. 206 - Reverse Linked List

## Core concept

It's really straightforward to reverse a linked list by simply reversing the
link/connection between each node in the list until the last node.

<p align="center"><img width="350"
src="https://media.geeksforgeeks.org/wp-content/cdn-uploads/RGIF2.gif" /></p>
<p align="center"><i>Source: GeeksForGeeks</i></p>

## Pseudocode

```text
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
FUNCTION reverseList(head):
  // Skip checking for empty list or single-element list
  IF head == nil OR head.Next == nil:
    RETURN head
  prev := nil
  cur := head
  WHILE cur != nil DO:
    next := cur.Next
    cur.Next = prev
    prev = cur
    cur = next
  ENDWHILE
  prev = cur // Assign back the last element before nil
  RETURN prev
```
