package kthlargestelementinanarray

import (
	"container/heap"
	"sort"
)

// findKthLargestBySorting finds the Kth largest element by sorting method
// Time complexity: O(n*logn)
// Space complexity: O(1)
func findKthLargestBySorting(nums []int, k int) int {
	// Sort the numsay (in-place)
	sort.Ints(nums)
	// Kth largest element can be found by len(nums) - k
	return nums[len(nums)-k]
}

/* * Max Heap * */
type IntHeap []int

func (h IntHeap) Len() int               { return len(h) }
func (h IntHeap) Less(i int, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// findKthLargestByHeap finds the Kth largest element using Heap
// Tine complexity: O(n*logn)
// Space complexity: O(n) - Auxiliary array for heap storage
func findKthLargestByHeap(nums []int, k int) int {
	// Build a heap
	h := IntHeap(nums)
	heap.Init(&h)
	res := 0
	for k > 0 {
		res = heap.Pop(&h).(int)
		k--
	}
	return res
}

// findKthLargestByQuickSelect finds Kth largest element by Quick Select algorithm
// Time complexity: O(2n) - Worst: O(n^2)
// Space Complexity: O(logn) (due to call stack)
func findKthLargestByQuickSelect(nums []int, k int) int {
	k = len(nums) - k
	ans := quickSelect(nums, k, 0, len(nums)-1)
	return nums[ans]
}

/* * Helper methods * */
// quickSelect selects the pivot that's Kth largest/smallest element in the array
func quickSelect(nums []int, k int, l int, r int) int {
	// Partition and get pivot index
	pi := partition(nums, l, r)
	if pi > k {
		return quickSelect(nums, k, l, pi-1)
	} else if pi < k {
		return quickSelect(nums, k, pi+1, r)
	}
	// If found the index
	return pi
}

// Partition partitions the array using pivot element
func partition(nums []int, l int, r int) int {
	pivot := nums[r]
	p := l
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[p], nums[j] = nums[j], nums[p]
			p++
		}
	}
	nums[r], nums[p] = nums[p], nums[r]
	return p
}
