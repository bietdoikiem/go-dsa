package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubble(t *testing.T) {
	assert := assert.New(t)
	// BubbleSort
	arr := []int{7, 5, 4, 2}
	BubbleSort(arr)
	// Displays the result
	Display(arr)
	// Assertion
	assert.Equal(7, arr[3], "Last element should be 7.")
	assert.Equal(2, arr[0], "First element should be 2.")
}
