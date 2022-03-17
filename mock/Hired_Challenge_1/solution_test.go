package hiredchallenge1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution_1(t *testing.T) {
	assert := assert.New(t)
	ans := Solution([]int64{3, 6, 2, 9, -1, 10})
	assert.Equal("Left", ans)
}

func Test_Solution_2(t *testing.T) {
	assert := assert.New(t)
	ans := Solution([]int64{1, 10, 5, 1, 0, 6})
	assert.Equal("", ans)
}
