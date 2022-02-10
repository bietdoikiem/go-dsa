<p align="center"><img width="550" src="https://i.ibb.co/tKHKT9J/image.png" /></p>

# No. 215 - Kth Largest Element in an Array

## Core Concept

In this problem, we're going to solve the problems in 3 ways:

### 1. Sorts the array

Since we're asked to find the Kth largest element, we can definitely sort the
array in ascending order and obtain the Kth largest one by retrieving the
element at (_array's length_ - k)<sup>th</sup> index

<p align="center"><img
src="https://i.ibb.co/5nnPJQM/Screenshot-20220210-144438-Samsung-Notes.jpg"
width="350" /></p>

**Note:**

> This problem is heavily relied on sorting algorithms. The most time-efficient
> sorting algorithsm are, for example, Quick Sort and Merge Sort. We're going to
> use the default Sorting algorithm in Go's standard package which is Quick Sort
> in the code.

### 2. Max Heap

As we're considering for the Kth largest element, a Max Heap data structure would
fit perfectly for the case. Each of the element in the Max Heap are stored such
that the parent node is always larger than one/two of its children. Thus, we can
retrieve the Kth largest element by popping the k number of elements from the
Heap.

<p align="center"><img
src="https://i.ibb.co/rb3KZnq/Screenshot-20220210-145812-Samsung-Notes.jpg"
width="350" /></p>

### 3. Quick Select

Quick Select algorithm is pretty much like the Quick Sort algorithm, in which we
need to partition the array in a way that the there are n elements lower and m
elements higher from the pivot's index.

However, there's a difference between Quick Sort and Quick Select is that we're not trying to sort the whole array,
instead, we're going to use partition technique to narrow down the selection
range until the pivot index is equal to (_array's length_ - k)<sup>th</sup> and
that's the Kth Largest element!

<p align="center"><img
src="https://i.ibb.co/47XbBzd/20220210-151228.jpg"
width="350" /></p>

## Pseudocode

### 1. Sorts the array

Time complexity: O(n\*logn)

Space complexity: O(1)

```text
FUNCTION findKthLargest(nums, k):
  kIndex := length(nums) - k
  nums.Sort() // Sorts the array using Quick Sort
  return nums[kIndex]
```

### 2. Max Heap

Time complexity: O(n*logh + k*logn) // O(n*logh) to build a Heap & O(k*logn) to pop
k elements

Space Complexity: O(h)

```text
FUNCTION findKthLargest(nums, k):
  h := heap.Build(nums) // Build a heap from nums array
  res := 0
  WHILE k > 0:
    res = Pop(h) // Pop the highest element from the heap
    k--
  return res
```

### 3. Quick Select

Time Complexity: O(n) Average Case - O(n<sup>2</sup>) Worst case in which array
is already sorted -> Can be mitigated using random pivot index

Space Complexity: O(1)

```text
FUNCTION findKthLargest(nums, k):
  pi := quickSelect(nums, k, 0, length(k) - 1)
  RETURN pi

FUNCTION quickSelect(nums, k, l, r):
  pi := partition(nums, l, r)
  IF pi > k:
    RETURN quickSelect(nums, l, pi-1)
  ELSE pi < k:
    RETURN quickSelect(nums, pi+1, r)
  RETURN pi

FUNCTION partition(nums, l, r):
  pivot := nums[r] // Right-most element
  pi := l // Swapped pivot index
  FOR j := l TO r DO:
    IF arr[j] < pivot:
      Swap(arr, pi, j)
      pi++
  ENDFOR
  Swap(arr, pi, r)
  RETURN pi
```
