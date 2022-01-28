package trees

import (
	"constraints"
	"fmt"
)

func DisplayBSTInOrder[T constraints.Ordered](root *BSTNode[T]) {
	if root == nil {
		return
	}
	DisplayBSTInOrder(root.left)
	fmt.Print(root.value, " ")
	DisplayBSTInOrder(root.right)
}
