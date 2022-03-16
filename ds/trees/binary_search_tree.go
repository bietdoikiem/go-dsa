package trees

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// TODO: Methods for Binary Search Tree
// 1. Insert ✅
// 2. Search ✅
// 3. Min ✅
// 4. Max ✅
// 5. Delete (key)

type BSTNode[T constraints.Ordered] struct {
	value T
	left  *BSTNode[T]
	right *BSTNode[T]
}

// NewBSTNode instantiates a new node for the Binary Search Tree
func NewBSTNode[T constraints.Ordered](value T) BSTNode[T] {
	return BSTNode[T]{value: value}
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *BSTNode[T]
}

// NewBinarySearchTree instantiates a Binary Search Tree
func NewBinarySearchTree[T constraints.Ordered]() BinarySearchTree[T] {
	return BinarySearchTree[T]{}
}

// Insert inserts a new node to the tree
func (t *BinarySearchTree[T]) Insert(value T) {
	t.root = recursiveInsert(t.root, value)
}

// recursiveInsert inserts a new node to the tree recursively
func recursiveInsert[T constraints.Ordered](root *BSTNode[T], value T) *BSTNode[T] {
	if root == nil {
		node := NewBSTNode(value)
		return &node
	}
	// Traverse recursively
	if value < root.value {
		root.left = recursiveInsert(root.left, value)
	} else {
		root.right = recursiveInsert(root.right, value)
	}
	// Return unchanged node
	return root
}

// Search searches for a value in the BST
func (t *BinarySearchTree[T]) Search(value T) (*T, error) {
	v := recursiveSearch(t.root, value)
	if v == nil {
		return nil, fmt.Errorf("binarysearchtree error: cannot search for value %v", value)
	}
	return v, nil
}

// recursiveSearch searches recursively for a value in the BST
func recursiveSearch[T constraints.Ordered](root *BSTNode[T], value T) *T {
	if root == nil {
		return nil
	}
	// Compare cases
	if value < root.value {
		return recursiveSearch(root.left, value)
	}
	if value > root.value {
		return recursiveSearch(root.right, value)
	}
	// If found
	return &root.value
}

// Min finds the minimum value in the given BST tree
func (t *BinarySearchTree[T]) Min() (*T, error) {
	// Check empty tree
	if t.root == nil {
		return nil, errors.New("binarysearchtree error: the tree is empty")
	}
	min := recursiveMin(t.root)
	return &min, nil
}

// recursiveMin recursively finds the minimum value of the BST
func recursiveMin[T constraints.Ordered](root *BSTNode[T]) T {
	min := root.value
	for root.left != nil {
		min = root.left.value
		root = root.left
	}
	return min
}

// Max returns the maximum value of the BST Tree
func (t *BinarySearchTree[T]) Max() (*T, error) {
	if t.IsEmpty() {
		return nil, errors.New("binarysearchtree error: the tree is empty")
	}
	max := recursiveMax(t.root)
	return &max, nil
}

// recursiveMax recursively finds obtain the maximum value of the BST
func recursiveMax[T constraints.Ordered](root *BSTNode[T]) T {
	max := root.value
	for root.right != nil {
		max = root.right.value
		root = root.right
	}
	return max
}

// Delete deletes a node from the BST
func (t *BinarySearchTree[T]) Delete(value T) (bool, error) {
	deleteFlag := false
	t.root = recursiveDelete(t.root, value, &deleteFlag)
	if !deleteFlag {
		return false, fmt.Errorf("binarysearchtreeerror: cannot delete as value %v does not exist", value)
	}
	return true, nil
}

// recursiveDelete recursively deletes the given value/key from the BST
func recursiveDelete[T constraints.Ordered](root *BSTNode[T], value T, deleteSuccess *bool) *BSTNode[T] {
	// Base case: Tree is empty
	if root == nil {
		return nil
	}
	// 2nd Case: Node deleted has one child
	// Check deleted value
	if value < root.value {
		root.left = recursiveDelete(root.left, value, deleteSuccess)
	} else if value > root.value {
		root.right = recursiveDelete(root.right, value, deleteSuccess)
	} else {
		// If found the deleted key
		*deleteSuccess = true
		// Node with only one child (left or right)
		// Delete by replacing it with its child
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}
		// Node with two children
		// Replace root with the minimum successor in the right subtree OR
		// Replace root with the maximum successor in the left subtree
		root.value = recursiveMax(root.left)
		// Delete the in-order left successor
		root.left = recursiveDelete(root.left, root.value, deleteSuccess)
	}
	// Return the unchanged node
	return root
}

// IsEmpty checks if the tree is currently empty (without parent root)
func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.root == nil
}
