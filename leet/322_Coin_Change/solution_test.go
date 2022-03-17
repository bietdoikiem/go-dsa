package coinchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

+Example 2:
Input: coins = [2], amount = 3
Output: -1

+Example 3:
Input: coins = [1], amount = 0
Output: 0

+Example 4:
Input: coins = [1,2,5], amount = 100
Output: 20

+Example 5:
Input: [186,419,83,408], amount = 6249
Output: ?
*/

func Test_coinChange_1(t *testing.T) {
	assert := assert.New(t)
	ans := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(3, ans)
}

func Test_coinChange_2(t *testing.T) {
	assert := assert.New(t)
	ans := coinChange([]int{2}, 3)
	assert.Equal(-1, ans)
}

func Test_coinChange_3(t *testing.T) {
	assert := assert.New(t)
	ans := coinChange([]int{1}, 0)
	assert.Equal(0, ans)
}

func Test_coinChange_4(t *testing.T) {
	assert := assert.New(t)
	ans := coinChange([]int{1, 2, 5}, 100)
	assert.Equal(20, ans)
}

func Test_coinChange_5(t *testing.T) {
	// assert := assert.New(t)
	ans := coinChange([]int{186, 419, 83, 408}, 6249)
	fmt.Println("Answer is:", ans)
	// assert.Equal(20, ans)
}

func Test_coinChangeBottomUp(t *testing.T) {
	assert := assert.New(t)
	ans := coinChangeBottomUp([]int{1, 2, 5}, 100)
	assert.Equal(20, ans)
}
