package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
+Example 1:
Input: n = 3
Output: 3
*/

func Test_climbStairs_1(t *testing.T) {
	assert := assert.New(t)
	ans := climbStairs(3)
	assert.Equal(3, ans, "There should be 3 possible ways to climb the stair.")
}
