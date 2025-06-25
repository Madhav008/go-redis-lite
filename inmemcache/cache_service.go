// cache_service.go
package inmemcache

import (
	"fmt"
	"sync"
	"time"
)

type CachedData struct {
	value string
	ttl   int
}


type CacheService struct {
	data  map[string]CachedData
	mutex sync.Mutex
	quit  chan struct{}
}

func NewCacheService() *CacheService {
	return &CacheService{
		data: make(map[string]CachedData),
		quit: make(chan struct{}),
	}
}

func (c *CacheService) Start() error {
	fmt.Println("Cache service is starting...")
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				c.UpdateTtl()
			case <-c.quit:
				return
			}
		}
	}()
	return nil
}

func (c *CacheService) UpdateTtl() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, val := range c.data {
		val.ttl--
		if val.ttl <= 0 {
			delete(c.data, key)
			continue
		}
		c.data[key] = val
	}
}

func (c *CacheService) Stop() error {
	fmt.Println("Cache service is stopping...")
	close(c.quit)
	return nil
}

func (c *CacheService) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = CachedData{value: value, ttl: 60}
}

func (c *CacheService) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.data[key]
	return val.value, ok
}

func (c *CacheService) Update(key, value string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.data[key]; ok {
		c.data[key] = CachedData{value: value, ttl: 60}
		return value, true
	}
	return "", false
}

func (c *CacheService) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}
