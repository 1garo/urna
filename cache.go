package urna

import (
	"container/heap"
	"sync"
)

/*
1. set(key, value) to create or update a key-value pair
2. get(key) to return a value for a given key
3. delete(key) to hard delete a particular value pair
4. clear() to clear all data from cache.
5. Implement evicion strategy (cache replacement policies)
*/

type MinHeap []*item
// An MinHeap is a min-heap of ints.

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*item))
}

func (h *MinHeap) Pop() *item {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type cache struct {
	size int
	mu sync.RWMutex
	data map[string]*item

	minHeap *MinHeap
}

type item struct {
	key string
	value string
	frequency int
	index int
}

// NewCache create a new *cache type
func NewCache() *cache {
	var a MinHeap
	var d map[string]*item
	t := a.Pop()
	delete(d, t.key)

	return &cache {
		data: make(map[string]*item, 0),
		mu: sync.RWMutex{},
	}
}

// Get get value from cache
func (c *cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok 
}

// GetMultiple get multiple values from cache
func (c *cache) GetMultiple(keys []string) []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	results := make([]string, 0)

	for _, key := range keys {
		if val, ok := c.data[key]; ok {
			results = append(results, val)
		}
	}

	return results
}

// Set set a value into the cache
func (c *cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// SetMultiple set multiple values into cache
func (c *cache) SetMultiple(items map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range items {
		c.data[k] = v
	}
}

// Clear clear all cache
func (c *cache) Clear() {
	clear(c.data)
}

// Delete delete key from cache
func (c *cache) Delete(key string) {
	delete(c.data, key)
}

// Len retrieve the length of cache data
func (c *cache) Len() int {
	return len(c.data)
}

