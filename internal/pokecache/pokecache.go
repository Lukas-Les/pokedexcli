package pokecache

import (
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntry map[string]CacheEntry
	interval   time.Duration
}

func NewCache(interval time.Duration) Cache {
	c := Cache{cacheEntry: make(map[string]CacheEntry), interval: interval}
	go c.ReapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	_, exists := c.cacheEntry[key]
	if exists {
		return
	}
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.cacheEntry[key]
	if !exists {
		return []byte{}, exists
	}
	return entry.val, exists
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.ReapOnce(time.Now())
	}
}

func (c *Cache) ReapOnce(now time.Time) {
	for key, entry := range c.cacheEntry {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.cacheEntry, key)
		}
	}
}
