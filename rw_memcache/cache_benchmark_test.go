package rw_memcache

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	cache := NewCacheService()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		val := "val" + strconv.Itoa(i)
		cache.Set(key, val)
	}
}

func BenchmarkGet(b *testing.B) {
	cache := NewCacheService()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Set(key, "value")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Get(key)
	}
}

func BenchmarkUpdate(b *testing.B) {
	cache := NewCacheService()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Set(key, "value")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Update(key, "new-value")
	}
}
