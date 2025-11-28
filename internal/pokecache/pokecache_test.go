package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	t.Run("Should add and get value from cache", func(t *testing.T) {
		cache := NewCache(time.Minute)
		val := []byte("something")
		cache.Add("test", val)
		fromCache, exists := cache.Get("test")
		if !exists {
			t.Errorf("%v was not added to cache", val)
		}
		if !bytes.Equal(fromCache, val) {
			t.Errorf("expected: %v, got: %v", val, fromCache)
		}
	})

	t.Run("Item from cache should get removed after time period", func(t *testing.T) {
		base := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
		cache := NewCache(time.Minute)
		cache.cacheEntry["old"] = CacheEntry{createdAt: base.Add(time.Minute * -2)}
		cache.cacheEntry["new"] = CacheEntry{createdAt: base.Add(time.Second * -1)}
		cache.ReapOnce(base)
		if _, ok := cache.cacheEntry["old"]; ok {
			t.Errorf("expected old to get deleted")
		}
		if _, ok := cache.cacheEntry["new"]; !ok {
			t.Errorf("expected new to remain")
		}
	})

}
