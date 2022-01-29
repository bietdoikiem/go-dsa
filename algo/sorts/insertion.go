package sorts

import (
	"constraints"
)

// InsertionSort sorts the elements of the array using Insertion Sort algorithm
// Time Complexity: O(n^2) - Stable
func InsertionSort[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i] // Current temp
		j := i - 1     // Inner loop counter
		for j >= 0 && temp < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = temp
	}
}
