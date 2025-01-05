package main

import (
	"fmt"
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

func demo() {
	// Create a new cache for string values with a TTL of 5 seconds
	cache := NewCache[string](5 * time.Second)

	// Add some data to the cache
	cache.Set("user:123", "John Doe")
	cache.Set("user:456", "Jane Smith")

	// Retrieve data from the cache
	if value, found := cache.Get("user:123"); found {
		fmt.Println("Cache hit for user:123:", value)
	} else {
		fmt.Println("Cache miss for user:123")
	}

	// Delete a key from the cache
	cache.Delete("user:456")

	// Check cache after deletion
	if value, found := cache.Get("user:456"); found {
		fmt.Println("Cache hit for user:456:", value)
	} else {
		fmt.Println("Cache miss for user:456")
	}

	// Wait to see if the TTL expires (example purposes)
	time.Sleep(6 * time.Second)

	// Check cache after TTL expiration
	if value, found := cache.Get("user:123"); found {
		fmt.Println("Cache hit for user:123:", value)
	} else {
		fmt.Println("Cache miss for user:123")
	}
}
