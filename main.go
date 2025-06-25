package main

import (
	"fmt"
	"go-redis-lite/models"
	"go-redis-lite/rw_memcache"
	"os"
	"os/signal"
	"syscall"
)

// App depends on the Cache interface, not a specific implementation
type App struct {
	cache models.Cache
}

// NewApp is a constructor that injects the dependency
func NewApp(cache models.Cache) *App {
	return &App{cache: cache}
}

func (a *App) Run() {
	a.cache.Set("username", "madhav")

	if val, ok := a.cache.Get("username"); ok {
		fmt.Println("Cached value:", val)
	}

	updated, ok := a.cache.Update("username", "jindal")
	if ok {
		fmt.Println("Updated value:", updated)
	}

	a.cache.Delete("username")
}

func main() {
	// You can switch this to a RedisCache that implements the same interface
	// inMemoryCache := inmemcache.NewCacheService()
	inMemoryCache := rw_memcache.NewCacheService()
	err := inMemoryCache.Start()
	if err != nil {
		fmt.Println("Failed to start cache service:", err)
		return
	}

	// Injecting the cache dependency into your app
	app := NewApp(inMemoryCache)
	app.Run()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	inMemoryCache.Stop()
	fmt.Println("Cache service stopped.")
}
