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

// RecursiveInsertionSort sorts the elements using Insertion Sort
// algorithm in recursion approach.
// Time complexity: still O(n^2) - Stable
func RecursiveInsertionSort[T constraints.Ordered](i int, arr []T) {
	// Base case: out of bounds index
	if i <= 0 {
		return
	}
	temp := arr[i]
	j := i - 1
	for j >= 0 && temp < arr[j] {
		RecursiveInsertionSort(i-1, arr)
		arr[j+1] = arr[j]
		j -= 1
	}
	arr[j+1] = temp
}
