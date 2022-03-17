package hiredchallenge1

import "fmt"

func Solution(arr []int64) string {
	n := int64(len(arr))
	if n == 0 {
		return ""
	}
	var parentIdx int64
	var leftSum int64
	helper(arr, 2*parentIdx+1, &leftSum)
	parentIdx = 0 // reset parentIdx
	var rightSum int64
	fmt.Println("coming to right")
	helper(arr, 2*parentIdx+2, &rightSum)
	fmt.Println("leftSum:", leftSum)
	fmt.Println("rightSum:", rightSum)
	// Check different cases for return result
	if leftSum == rightSum {
		return ""
	} else if leftSum > rightSum {
		return "Left"
	}
	return "Right"
}

func helper(arr []int64, curIdx int64, sum *int64) {
	// Base case
	if curIdx > int64(len(arr)) {
		return
	}
	// Aggregate result from two branches
	if arr[curIdx] != -1 {
		*sum += arr[curIdx]
	}
	fmt.Println("sum:", *sum, "- curIdx:", curIdx)
	if 2*curIdx+1 < int64(len(arr)) {
		helper(arr, 2*curIdx+1, sum)
	}
	if 2*curIdx+2 < int64(len(arr)) {
		helper(arr, 2*curIdx+2, sum)
	}
}
