package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binary_search_tree(t *testing.T) {
	assert := assert.New(t)
	// Insert
	tree := NewBinarySearchTree[int]()
	tree.Insert(100)
	tree.Insert(20)
	tree.Insert(500)
	tree.Insert(10)
	tree.Insert(30)
	fmt.Print("Tree after insertion: ")
	DisplayBSTInOrder(tree.root)
	fmt.Println("")
	assert.Equal(100, tree.root.value, "Tree root should be 100.")
	// Search
	s1, err := tree.Search(10)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(10, *s1, "Found value should be 10.")
	s2, err := tree.Search(500)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(500, *s2, "Found value should be 500.")
	_, err = tree.Search(6969)
	if err != nil {
		fmt.Println(err)
	}
	// Min
	min, err := tree.Min()
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(10, *min, "Minimum value should be 10.")
	// Max
	max, err := tree.Max()
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(500, *max, "Maximum value should 500.")
	// Delete
	_, err = tree.Delete(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Tree after deletion: ")
	DisplayBSTInOrder(tree.root)
}
