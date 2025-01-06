package main

import (
	"sync"
	"time"
)

type Cache[T any] struct {
	data map[string]T
	mu   sync.RWMutex
	ttl  time.Duration
}

func NewCache[T any](ttl time.Duration) *Cache[T] {
	return &Cache[T]{
		data: make(map[string]T),
		ttl:  ttl,
	}
}

// Set adds or updates a key-value pair in the cache
func (c *Cache[T]) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves a value by key from the cache
func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.data[key]
	return value, found
}

// Delete removes a key-value pair from the cache
func (c *Cache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Clear clears all cache entries
func (c *Cache[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]T)
}
