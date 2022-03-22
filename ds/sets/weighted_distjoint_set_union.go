package sets

type WeightedDisjointSetUnion[T any] struct {
	items        []T
	connectivity []int
	weight       []int
}

// NewWeightedDisjointSetUnion returns an instance of weighted distjoint subset with initialized array.
func NewWeightedDisjointSetUnion[T any](items []T) WeightedDisjointSetUnion[T] {
	n := len(items)
	// Init connection and weight (or rank)
	con := make([]int, n)
	w := make([]int, n)
	for i := 0; i < n; i++ {
		con[i] = i
		w[i] = 1
	}
	return WeightedDisjointSetUnion[T]{items: items, connectivity: con, weight: w}
}

// root finds the root element of the current one
// Time complexity: O(logn). Amortized cost: near constant (complementary with Union-by-Rank)
func (ds WeightedDisjointSetUnion[T]) root(i int) int {
	for ds.connectivity[i] != i {
		ds.connectivity[i] = ds.connectivity[ds.connectivity[i]] // Flatten till the grandparent node
		i = ds.connectivity[i]
	}
	return i
}

// Find finds union elements (with path compression)
func (ds WeightedDisjointSetUnion[T]) Find(i1 int, i2 int) bool {
	r1 := ds.root(i1)
	r2 := ds.root(i2)
	return r1 == r2
}

// Union unions two elements to the root of the other.
// Rank is just the upperbound to prevent skewed tree, not the actual depth (in path compression method)
func (ds *WeightedDisjointSetUnion[T]) Union(i1 int, i2 int) {
	r1 := ds.root(i1)
	r2 := ds.root(i2)
	// Check the weight
	if ds.weight[r1] < ds.weight[r2] {
		ds.connectivity[r1] = ds.connectivity[r2]
		ds.weight[r2] += ds.weight[r1]
	} else {
		ds.connectivity[r2] = ds.connectivity[r1]
		ds.weight[r1] += ds.weight[r2]
	}
}
