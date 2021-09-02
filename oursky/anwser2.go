package oursky

import (
	"container/list"
	"math"
	"time"
)

type kv struct {
	k            int
	v            int
	lastUsedTime int64
	weight       int
}

// Cache declare cache struct
type Cache struct {
	capacity int
	data     map[int]*list.Element
}

// NewCache ...
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		data:     make(map[int]*list.Element),
	}
}

// Get serve caller to get value by key
func (c *Cache) Get(key int) int {
	now := time.Now().UnixNano()

	if ele, ok := c.data[key]; ok {
		// update last used time
		c.data[key].Value = kv{
			k:            key,
			v:            ele.Value.(kv).v,
			lastUsedTime: now,
			weight:       ele.Value.(kv).weight,
		}

		return ele.Value.(kv).v
	}

	return -1
}

// Put serve caller to put key val into Cache
func (c *Cache) Put(key int, val int, weight int) {
	now := time.Now().UnixNano()

	// via get by key to update last used time
	c.Get(key)

	if len(c.data) == c.capacity {
		mini := math.MaxInt32
		var removeKey int

		for _, ele := range c.data {
			if now != ele.Value.(kv).lastUsedTime {
				w := int(float64(ele.Value.(kv).weight) / math.Log(float64(now-ele.Value.(kv).lastUsedTime+1)))
				if mini > w {
					mini = w
					removeKey = ele.Value.(kv).k
				}
			} else {
				w := ele.Value.(kv).weight / -100
				if mini > w {
					mini = w
					removeKey = ele.Value.(kv).k
				}
			}
		}

		delete(c.data, removeKey)
	}

	// insert key and value into the data
	c.data[key].Value = kv{
		k:            key,
		v:            val,
		lastUsedTime: now,
		weight:       weight,
	}
}
