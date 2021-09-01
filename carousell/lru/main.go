package main

import (
	"container/list"
	"fmt"
)

type kv struct {
	k int
	v int
}

// LRUCache ...
type LRUCache struct {
	capacity int
	data     map[int]*list.Element
	hits     *list.List
}

// NewLRUCache ...
func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		data:     make(map[int]*list.Element),
		hits:     list.New(),
	}
}

// Get ...
func (c *LRUCache) Get(k int) int {
	// get value by key
	if el, ok := c.data[k]; ok {
		// move the key to front in the hits linked list
		c.hits.MoveToFront(el)
		return el.Value.(kv).v
	}

	return -1
}

// Put ...
func (c *LRUCache) Put(k, v int) {
	// get the key and move the key to front in the hits
	if el, ok := c.data[k]; ok {
		el.Value = kv{k: k, v: v}
		c.hits.MoveToFront(el)
		return
	}

	// if data is full
	if c.hits.Len() == c.capacity {
		last := c.hits.Back()

		// remove last element
		delete(c.data, last.Value.(kv).k)
		c.hits.Remove(last)
	}

	// put kv
	c.data[k] = c.hits.PushFront(kv{k: k, v: v})
}

func main() {
	c := NewLRUCache(2)

	// return -1 because it hasn't any elements
	fmt.Println(c.Get(1))
	fmt.Println(c)

	// put kv with 1, 1
	c.Put(1, 1)
	fmt.Println(c)

	// put kv with 2, 2
	c.Put(2,2)
	fmt.Println(c)

	// get by key
	fmt.Println(c.Get(1))
	fmt.Println(c)

	// put kv with 3, 3 but capacity is full, so remove the key 2
	c.Put(3, 3)
	fmt.Println(c)
}
