package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entities map[string]cacheEntity
	mu       *sync.Mutex
	ttl      time.Duration
}

type cacheEntity struct {
	val       []byte
	createdAt time.Time
}

var cache Cache

const reapInterval time.Duration = 5 * time.Second

func NewCache(interval time.Duration) {
	cache = Cache{
		entities: make(map[string]cacheEntity),
		mu:       &sync.Mutex{},
		ttl:      interval,
	}

	ticker := time.NewTicker(reapInterval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				reap()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entities[key] = cacheEntity{
		val:       val,
		createdAt: time.Now(),
	}
}

func Get(key string) ([]byte, bool) {
	c, ok := cache.entities[key]
	if !ok {
		return []byte{}, false
	}

	return c.val, true
}

func reap() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	now := time.Now()

	for k, e := range cache.entities {
		if now.Sub(e.createdAt) > cache.ttl {
			delete(cache.entities, k)
		}
	}
}
