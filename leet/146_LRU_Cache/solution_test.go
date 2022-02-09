package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRUCache_1(t *testing.T) {
	assert := assert.New(t) // Assertion
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)     // cache is {1=1}
	lRUCache.Put(2, 2)     // cache is {1=1, 2=2}
	res := lRUCache.Get(1) // return 1
	assert.Equal(1, res, "Get value should be 1.")
	lRUCache.Put(3, 3)    // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	res = lRUCache.Get(2) // returns -1 (not found)
	assert.Equal(-1, res, "Get value should be -1 due to invalid key.")
	lRUCache.Put(4, 4)    // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	res = lRUCache.Get(1) // return -1 (not found)
	assert.Equal(-1, res, "Get value should be -1 due to invalid key.")
	res = lRUCache.Get(3) // return 3
	assert.Equal(3, res, "Get value should be 3.")
	res = lRUCache.Get(4) // return 4
	assert.Equal(4, res, "Get value should be 4.")
}

func Test_LRUCache_2(t *testing.T) {
	lRUCache := Constructor(1)
	lRUCache.Put(2, 1)
	lRUCache.Get(2)
	lRUCache.Put(3, 2)
	lRUCache.Get(2)
	lRUCache.Get(2)
}
