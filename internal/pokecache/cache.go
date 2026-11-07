package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]cacheEntry
	mu   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{Data: make(map[string]cacheEntry)}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.Data[key] = v
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.Data[key]
	if ok {
		return val.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.Data {
			if time.Since(val.createdAt) > interval {
				delete(c.Data, key)
			}
		}
		c.mu.Unlock()
	}
}
