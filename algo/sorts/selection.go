package sorts

import "constraints"

func SelectionSort[T constraints.Ordered](arr []T) {
	var min int
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			// Find min in the rest
			if arr[j] < arr[min] {
				min = j
			}
		}
		// Swap min to the current element
		Swap(arr, i, min)
	}
}
