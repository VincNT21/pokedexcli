package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// New Cache fonction
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

// Add method
func (c *Cache) Add(key string, value []byte) {
	// Manage Mutex lock
	c.mu.Lock()
	defer c.mu.Unlock()

	// Add the new cache Entry
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get method
func (c *Cache) Get(key string) ([]byte, bool) {
	// Manage Mutex lock for read
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Check if key exists and return value
	result, ok := c.cacheMap[key]
	if !ok {
		return []byte{}, false
	} else {
		return result.val, true
	}
}

// Reap Loop method
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for t := range ticker.C {
		c.mu.Lock()

		for key, val := range c.cacheMap {
			if t.Sub(val.createdAt) >= interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
