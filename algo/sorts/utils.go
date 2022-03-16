package sorts

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Swap swaps two elements based on their index
func Swap[T constraints.Ordered](arr []T, i1 int, i2 int) {
	temp := arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = temp
}

// Display prints all the elements from the array
func Display[T constraints.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}
