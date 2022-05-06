package go_cache

import (
	"go-cache/lru"
	"sync"
)

type cache struct {
	mu        sync.Mutex
	lru       *lru.Cache
	cacheByte int64
}

func (c *cache) add(key string, val ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheByte, nil)
	}
	c.lru.Add(key, val)
}

func (c *cache) get(key string) (v ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), true
	}
	return
}
