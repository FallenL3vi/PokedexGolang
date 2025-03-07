package pokecache

import (
	"sync"
	"time"
	"errors"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	data map[string]cacheEntry
	lock sync.Mutex
}

func NewCache(newInterval time.Duration) *Cache {
	cache := &Cache {
		data : make(map[string]cacheEntry),
	}

	go cache.reapLoop(newInterval)

	return cache
}

func (cache *Cache) Add(key string, new_val []byte) error{
	cache.lock.Lock()
	defer cache.lock.Unlock()
	if key == "" {
		return errors.New("Can not create cache entry with empty key")
	}

	if len(new_val) == 0 {
		return errors.New("Can not create cache entry with empty value")
	}

	//Maybe check if entry exists TO DO

	cache.data[key] = cacheEntry{createdAt : time.Now(), val : new_val}
	return nil

}


func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	if key == "" {
		return nil, false
	}
	
	if val, ok := cache.data[key]; ok {
		return val.val, ok
	}

	return nil, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case t := <-ticker.C:
			cache.lock.Lock()
			for key, val := range cache.data {
				if val.createdAt.Before(t) {
					delete(cache.data, key)
				}
			}
			cache.lock.Unlock()
		}
	}
}