package sets

type DisjointSet[T any] struct {
	items        []T
	connectivity []int
}

func NewDisjointSet[T any](items []T) DisjointSet[T] {
	// Fill connectivity array with items index
	con := make([]int, len(items))
	for i := 0; i < len(con); i++ {
		con[i] = i
	}
	return DisjointSet[T]{items, con}
}

// Union merges the two subsets altogether.
// Time complexity: O(n).
// CAUTION: There's a con in this function as if you have to union N elements in the set,
// it would take O(n^2) time complexity to finish.
func (ds *DisjointSet[T]) Union(i1, i2 int) {
	tmp := ds.connectivity[i1]
	for i := 0; i < len(ds.connectivity); i++ {
		if ds.connectivity[i] == tmp {
			ds.connectivity[i] = ds.connectivity[i2]
		}
	}
}

// Find finds and verifies if the two elements are connected with each other.
// Time complexity: O(n)
func (ds DisjointSet[T]) Find(i1, i2 int) bool {
	return ds.connectivity[i1] == ds.connectivity[i2]
}

/* Alternatives way to build a Disjoint Set */
/// Consider the set of each element is the root of itself.
