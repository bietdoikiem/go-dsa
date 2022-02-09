package lrucache

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*Node
	queue    LinkedQueue
}

// Node of queue
type Node struct {
	val  int
	key  int
	next *Node
	prev *Node
}

// Queue implemeneted in doubly linked list queue
type LinkedQueue struct {
	head     *Node
	tail     *Node
	capacity int
	size     int
}

// Constructor instantiates a new LRU Cache
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, cache: make(map[int]*Node), queue: LinkedQueue{capacity: capacity}}
}

// Get retrieves an item from the cache based on its key
// Time complexity: O(n)
func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// Only insert at the tail if the it's not already there
	if this.queue.tail != node {
		// Remove node from current position
		this.queue.Remove(node)
		// Insert at the end of recently resetted node
		this.queue.Enqueue(node)
	}
	return node.val
}

// Put puts a new item into the cache, evicts LRU item if cache is full
// Time complexity: O(n)
func (this *LRUCache) Put(key int, value int) {
	// Update existing key
	if existingNode, ok := this.cache[key]; ok {
		// Update current value
		existingNode.val = value
		// Move the node to the end of the queue if it's not in the tail already
		if this.queue.tail != existingNode {
			// Remove node from current position
			this.queue.Remove(existingNode)
			// Enqueue recent deleted node to the queue
			this.queue.Enqueue(existingNode)
		}
		return
	}
	// Insert new node to map
	inserted := &Node{key: key, val: value}
	if this.size < this.capacity {
		this.cache[key] = inserted
	} else {
		// Retrieve dequeued item from recent items
		leastUsed := this.queue.Dequeue()
		// Delete least used item from the map
		this.size -= 1
		delete(this.cache, leastUsed.key)
		// Add new item to cache
		this.cache[key] = inserted
	}
	this.size += 1
	this.queue.Enqueue(inserted) // Insert a new node
}

/* * Queue's Methods * */

// Dequeues dequeues the first element of the queue
func (this *LinkedQueue) Dequeue() *Node {
	// Dequeue empty list
	if this.size == 0 {
		return nil
	}
	dequeued := this.head // Dequeue first item (FIFO)
	// Dequeue single-element queue
	if this.size == 1 {
		this.head = nil
		this.tail = nil
		this.size -= 1
		return dequeued
	}
	// Dequeue more-than-one element queue
	this.head = this.head.next
	// Current head is null
	this.head.prev = nil
	this.size -= 1
	// Check if it's the only remaning element after dequeuing
	if this.size == 1 {
		this.tail = this.head
	}
	return dequeued
}

// Remove removes a node from the linked queue
func (this *LinkedQueue) Remove(node *Node) {
	if this.size == 0 && this.head == nil {
		return
	}
	// Remove at head only
	if this.head == node && this.size == 1 {
		this.head = nil
		this.tail = nil
		this.size -= 1
		return
	}
	// Remove at head
	if this.head == node {
		next := this.head.next
		next.prev = nil
		this.head.next = nil
		this.head = next
		this.size -= 1
		// Set last element to be head & tail if this is the only left one
		if this.size == 1 {
			this.tail = this.head
		}
		// Set next element to be tail if there only 2 remains
		if this.size == 2 {
			this.tail = this.head.next
		}
		return
	}
	// Remove at tail
	if this.tail == node {
		prev := this.tail.prev
		// Cut the link between prev and current tail
		prev.next = nil
		this.tail.prev = nil
		this.tail = prev
		this.size -= 1
		// Set last element to be head & tail if this is the only left one
		if this.size == 1 {
			this.tail = this.head
		}
		return
	}
	// Remove at body
	node.prev.next = node.next
	node.next.prev = node.prev
}

// Enqueue appends a new node to the tail of linked queue
func (this *LinkedQueue) Enqueue(node *Node) {
	// Insert at empty list
	if this.size == 0 {
		this.head = node
		this.tail = node
	} else {
		// Insert at tail
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.size += 1
}
