package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	CacheMap map[string]cacheEntry
	Mu       sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	// Create the new cacheEntry
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	// Manage Mutex lock
	c.Mu.Lock()
	defer c.Mu.Unlock()

	// Add the new cache Entry
	c.CacheMap[key] = newEntry
}

func (c *Cache) Get()

func NewCache(interval time.Duration) {

}
