package pokecache

import (
	"sync"
	"time"
)

func CreateNewCache(interval time.Duration) Cache {
	c := Cache{
		vals: make(map[string]cacheEntry),
		mutex:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}
 
func (c *Cache) Add (key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.vals[key] = cacheEntry{
		createdAt: time.Now().Unix(),
		val: val,
	}
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.vals[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)

		c.mutex.Lock()
		for key, entry := range c.vals {
			currentTime := time.Now().Unix()
			if entry.createdAt < currentTime - int64(interval.Seconds()) {
				delete(c.vals, key)
			}
		}
		c.mutex.Unlock()
	}
}