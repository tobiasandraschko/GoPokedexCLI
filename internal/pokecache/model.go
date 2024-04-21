package pokecache

import "sync"

type Cache struct {
	vals map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt int64
	val       []byte
}
