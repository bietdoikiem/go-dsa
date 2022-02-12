package designhashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashMap_1(t *testing.T) {
	assert := assert.New(t) // Assertion
	obj := Constructor()
	obj.Put(1, 1)
	getItem := obj.Get(1)
	assert.Equal(1, getItem, "Retrieved item's value should be 1.")
	obj.Remove(1)
	anotherGet := obj.Get(1)
	assert.Equal(-1, anotherGet, "Retrieved item's should be -1 as it has been removed already.")
}
