package scl2022

import (
	"fmt"
	"testing"
)

/*
+Example 1:
Input: [1,2,3,6]
Output: 6
*/

func Test_Solution_1(t *testing.T) {
	ans := findMaxDisjointSubset([]int{1, 2, 3, 6})
	fmt.Println("ans:", ans)
}
