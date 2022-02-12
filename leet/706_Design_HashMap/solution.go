package designhashmap

import "container/list"

// Not an optimal solution (considering to be optimal if have Load Factor)

const Capacity = 89 // Prime number xD

type HashNode struct {
	key   int
	value int
}

type MyHashMap struct {
	elements [Capacity]*list.List
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

// Put puts a key-value pair into the hash map
func (this *MyHashMap) Put(key int, value int) {
	// Cases to consider upon adding a new value
	index := hash(key)
	// 1. Key already exists, consider updating the value
	if this.Get(key) != -1 {
		for e := this.elements[index].Front(); e != nil; e = e.Next() {
			if e.Value.(HashNode).key == key {
				e.Value = HashNode{key: key, value: value}
				return
			}
		}
	}
	// 2. Put a completely new value in to the uninitialized index
	if this.elements[index] == nil {
		// Init a new list
		this.elements[index] = list.New()
		// Push to that list
		this.elements[index].PushBack(HashNode{key, value})
	} else {
		// 3. There's a collision -> Append new key value to the list
		this.elements[index].PushBack(HashNode{key, value})
	}
}

func (this *MyHashMap) Get(key int) int {
	// Cases to consider when get new map
	index := hash(key)
	// Empty/Uninitialized index
	if this.elements[index] == nil {
		return -1
	}
	// Check if key-pair exist in the list
	for e := this.elements[index].Front(); e != nil; e = e.Next() {
		if e.Value.(HashNode).key == key {
			return e.Value.(HashNode).value
		}
	}
	return -1
}

// Remove removes a key-value pair from the map
func (this *MyHashMap) Remove(key int) {
	// Cases to consider
	// 1. Key not exist
	if this.Get(key) == -1 {
		return
	}
	index := hash(key)
	// 2a. Key exist at only-1 index -> Dereference it
	if this.elements[index].Len() == 1 {
		this.elements[index] = nil
		return
	}
	// 2b. Key exist at certain index
	for e := this.elements[index].Front(); e != nil; e = e.Next() {
		// If match key -> delete
		if e.Value.(HashNode).key == key {
			this.elements[index].Remove(e)
		}
	}
}

func hash(key int) int {
	return key % Capacity // The power of modulo xD
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
