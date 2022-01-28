package sorts

import "constraints"

// BubbleSort sorts the array using Bubble Sort algorithm
// Time complexity: O(n^2) - Not Stable
func BubbleSort[T constraints.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
}
