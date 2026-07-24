package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	log.Println("Searching in cache for data")
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]

	if !ok {
		return nil, false
	}

	return append([]byte{}, entry.val...), true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				log.Println("CACHE cleaned!", string(key))
				delete(c.entries, key)
			}
		}

		c.mu.Unlock()
	}
}
