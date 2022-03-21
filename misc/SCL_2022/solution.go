package scl2022

import "fmt"

func findMaxDisjointSubset(arr []int) int {
	subsetSum := []int{}
	for i := 0; i < len(arr); i++ {
		helper(arr, i, 0, 0, &subsetSum)
	}
	fmt.Println("subsetSum:", subsetSum)
	var maxPairVal int
	for i := 0; i < len(subsetSum); i++ {
		for j := i + 1; j < len(subsetSum); j++ {
			if subsetSum[i] == subsetSum[j] {
				maxPairVal = max(maxPairVal, subsetSum[i])
			}
		}
	}
	return maxPairVal
}

func helper(arr []int, curIdx int, curLen int, curSum int, subsetSum *[]int) {
	if curIdx >= len(arr) {
		return
	}
	// Subset should NOT equal or exceed max length
	if curLen >= len(arr)-1 {
		return
	}
	curSum += arr[curIdx]
	curLen += 1
	*subsetSum = append(*subsetSum, curSum)
	helper(arr, curIdx+1, curLen, curSum, subsetSum)
	helper(arr, curIdx+2, curLen, curSum, subsetSum)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
