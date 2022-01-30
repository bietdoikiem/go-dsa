package sorts

import (
	"constraints"
)

// MergeSort sorts elements of arrrays (ascending order) using Merge Sort algorithm
// Time complexity: O(n*log(n)) - Stable
func MergeSort[T constraints.Ordered](arr []T, optionals ...bool) []T {
	// Check empty
	if len(arr) == 0 {
		return arr
	}
	var inPlace bool
	if len(optionals) > 0 {
		inPlace = optionals[0]
	}
	sorted := recursiveMergeSort(arr)
	// In-place copy (Cost: O(n))
	if inPlace {
		for i := 0; i < len(arr); i++ {
			arr[i] = sorted[i]
		}
		return arr
	}
	return sorted
}

// recursiveMergeSort (helper function) sorts recursively using Merge Sort algorithm
func recursiveMergeSort[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	if n == 1 {
		return arr[:1]
	}
	m := int(n / 2)
	merged := merge(recursiveMergeSort(arr[0:m]), recursiveMergeSort(arr[m:n]))
	return merged
}

// Merge two arrays (ascending)
func merge[T constraints.Ordered](a1 []T, a2 []T) []T {
	// Length
	n1 := len(a1)
	n2 := len(a2)
	merged := make([]T, n1+n2)
	// Pointers
	i := 0
	j := 0
	k := 0
	// Merge
	for i < n1 && j < n2 {
		if a1[i] <= a2[j] {
			// Add a1 items to the merged array
			merged[k] = a1[i]
			i += 1
		} else {
			// Add a2 items to the merged array
			merged[k] = a2[j]
			j += 1
		}
		k += 1 // Update k pointer after insertion
	}
	// Copy remaining elements of a1 if there are ones
	for i < n1 {
		merged[k] = a1[i]
		i += 1
		k += 1
	}
	// Else for a2
	for j < n2 {
		merged[k] = a2[j]
		j += 1
		k += 1
	}
	return merged
}
