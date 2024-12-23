package cache

import (
	"errors"
	"sync"
)

// Cache represents a generic cache with variable number of keys.
type Cache[K comparable, V any] struct {
	data map[K]any
	mu   sync.RWMutex
}

// NewCache creates a new generic cache.
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]any),
	}
}

// Get retrieves a value from the cache using a slice of keys.
func (c *Cache[K, V]) Get(keys []K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var zeroValue V
	node := any(c.data)
	for _, key := range keys {
		if nestedMap, ok := node.(map[K]any); ok {
			if nextNode, exists := nestedMap[key]; exists {
				node = nextNode
			} else {
				return zeroValue, false
			}
		} else {
			return zeroValue, false
		}
	}

	// At the end, node should be the value.
	if value, ok := node.(V); ok {
		return value, true
	}
	return zeroValue, false
}

// Set adds or updates a value in the cache using a slice of keys.
func (c *Cache[K, V]) Set(keys []K, value V) error {
	if len(keys) == 0 {
		return errors.New("keys cannot be empty")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	node := c.data
	for i, key := range keys {
		if i == len(keys)-1 {
			// Final key, store the value
			node[key] = value
		} else {
			// Intermediate key, ensure the next level exists
			if _, exists := node[key]; !exists {
				node[key] = make(map[K]any)
			}
			node = node[key].(map[K]any)
		}
	}

	return nil
}

// Reset clears the entire cache.
func (c *Cache[K, V]) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[K]any)
}
