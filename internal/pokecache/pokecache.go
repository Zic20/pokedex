package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}

	go c.reapLoop(interval)
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}

}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries,k)
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
