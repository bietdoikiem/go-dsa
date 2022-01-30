package sorts

import "constraints"

func QuickSort[T constraints.Ordered](arr []T) {
	// Check empty
	if len(arr) == 0 {
		return
	}
	recursiveQuickSort(arr, 0, len(arr)-1)
}

// recursiveQuickSort recursively sorts the elements using QuickSort algorithm
func recursiveQuickSort[T constraints.Ordered](arr []T, l int, r int) {
	// Base case
	if l >= r {
		return
	}
	pi := partition(arr, l, r)
	recursiveQuickSort(arr, l, pi-1)
	recursiveQuickSort(arr, pi+1, r)
}

// partition partitions the array into two parts: one whose values are smaller than
// the pivot, and the other contains larger values. It also returns the index of the pivot after partitioning
func partition[T constraints.Ordered](arr []T, l int, r int) int {
	pivot := arr[r] // Rightmost element as the pivot
	i := l
	// Partition into two parts
	for j := l; j < r; j++ {
		if arr[j] < pivot {
			Swap(arr, i, j)
			i += 1
		}
	}
	// Switch the last larger value
	Swap(arr, i, r)
	return i
}
