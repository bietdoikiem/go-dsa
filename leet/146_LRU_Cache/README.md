<p align="center"><img width="350" src="https://i.ibb.co/ysfNyHV/image.png" /></p>

# No. 146 - LRU Cache

## Core concept

Since the requirement of the problem indicates constant time O(1) for Get and
Put operation, we can immediately relate to the HashTable or HashMap.

Least Recently Used Item can be keep track using FIFO functionality (which is a
queue) since the first item to be added can be safely dequeued/staled from
the cache if it's not been renew or retrieved. Moreover, for this problem, we're
going to implement the queue using Doubly Linked List as its complexity for
removing at the front (enqueue) and inserting at the back (dequeue) are constant
O(1).

However, linked list structure does not allow us to get constant-time access to
the list of elements, which may prevent us from renewing an element in the
middle of the list. Luckily, we can take advantage of the previously created
HashMap of the cache to store value as a pointer to the linked-list queue in and
gain constant-time access pretty easily!

<p align="center"><img width="350"
src="https://i.ibb.co/D965Wn9/Screenshot-20220209-160604-Samsung-Notes.jpg" /></p>

## Pseudocode

```text
/* * Structs/Objects defined for the problem * */
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

FUNCTION Get(key):
  If key not exists in cache:
    RETURN -1
  node := cache[key]
  // Check if the key is not in the tail of the list already
  IF queue.tail != node:
    // Remove from the queue
    queue.Remove(node)
    // Renew by enqueuing to the tail of the linked-list queue
    queue.Enqueue(node)
  RETURN node

FUNCTION Put(key, value):
  IF key exists in the cache:
    cache[key] = value // Update existing one
    // Renew the queue if it's not already in the tail
    IF queue.tail != cache[key]:
      queue.Remove(node)
      queue.Enqueue(node)
    RETURN
  // If key not exists
  inserted := Node(key, value)
  IF LRUCache.size < LRUCache.capacity:
    cache[key] = inserted
  ELSE:
    // Retrieve least recent used item (dequeued one)
    least := queue.Dequeue()
    // Remove from hash map
    LRUCache.size -= 1
    delete(cache, least.key)
    // Insert new one
    cache[key] = inserted
  // Update queue and cache size
  LRUCache.size += 1
  queue.Enqueue(inserted)
```
