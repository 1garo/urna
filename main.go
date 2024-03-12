package main

import (
	"fmt"
	"sync"
)



/*
1. put(key, value) to create or update a key-value pair
2. get(key) to return a value for a given key
3. delete(key) to hard delete a particular value pair
4. clear() to clear all data from cache.
5. Implement evicion strategy (cache replacement policies)
*/
type cache struct {
	mu sync.RWMutex
	data map[string]string
}

func NewCache() *cache {
	return &cache {
		data: make(map[string]string, 0),
		mu: sync.RWMutex{},
	}
}

func (c *cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok 
}

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

func (c *cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *cache) SetMultiple(items map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range items {
		c.data[k] = v
	}
}

func (c *cache) Clear() {
	clear(c.data)
}

func (c *cache) Delete(key string) {
	delete(c.data, key)
}

func (c *cache) Len() int {
	return len(c.data)
}

func main() {
	fmt.Println("hello world")
}
