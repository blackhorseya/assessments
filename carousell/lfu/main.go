package main

import (
	"container/list"
	"fmt"
	"math"
	"time"
)

type kv struct {
	k            int
	v            int
	lastUsedTime int64
	freq         int
}

type LFUCache struct {
	capacity int
	data     map[int]*list.Element
	hits     *list.List
}

func New(cap int) *LFUCache {
	return &LFUCache{
		capacity: cap,
		data:     make(map[int]*list.Element),
		hits:     list.New(),
	}
}

func (c *LFUCache) Get(key int) int {
	// get the key, I must add freq 1 and update lastUsedTime for now
	if el, ok := c.data[key]; ok {
		// add 1 to freq
		c.data[key].Value = kv{
			k:            key,
			v:            el.Value.(kv).v,
			freq:         el.Value.(kv).freq + 1,
			lastUsedTime: time.Now().UnixNano(),
		}

		return el.Value.(kv).v
	}

	return -1
}

func (c *LFUCache) Put(k, v int) {
	now := time.Now().UnixNano()
	newEle := kv{
		k:            k,
		v:            v,
		freq:         1,
		lastUsedTime: now,
	}

	// put the key and value, I must add freq 1 and update lastUsedTime
	if el, ok := c.data[k]; ok {
		c.data[el.Value.(kv).k].Value = kv{
			k:            k,
			v:            v,
			freq:         el.Value.(kv).freq + 1,
			lastUsedTime: now,
		}
		return
	}

	// if the data is full, so I must remove less used element
	// I need to find the element of less freq and remove it.
	if len(c.data) == c.capacity {
		mini := math.MaxInt32
		minT := now
		var key int

		// find the less used or less time and store the key
		for _, el := range c.data {
			if mini > el.Value.(kv).freq {
				mini = el.Value.(kv).freq
				minT = el.Value.(kv).lastUsedTime
				key = el.Value.(kv).k
			} else if mini == el.Value.(kv).freq && minT > el.Value.(kv).lastUsedTime {
				key = el.Value.(kv).k
				minT = el.Value.(kv).lastUsedTime
			}
		}

		delete(c.data, c.data[key].Value.(kv).k)
	}

	// put the key and value into the data
	c.data[k] = c.hits.PushFront(newEle)
}

func main() {
	c := New(5)

	fmt.Println(c.Get(1))
	fmt.Println(c)

	c.Put(1, 1)
	fmt.Println(c)

	c.Put(2, 2)
	fmt.Println(c)

	c.Put(3, 3)
	fmt.Println(c)

	c.Put(3, 3)
	fmt.Println(c)

	c.Put(4, 4)
	fmt.Println(c)

	c.Put(4, 4)
	fmt.Println(c)

	c.Put(5, 5)
	fmt.Println(c)

	c.Put(1, 1)
	fmt.Println(c)

	c.Put(6, 6)
	fmt.Println(c)
}
