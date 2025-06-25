## 🚀 Built My Own Redis-Like In-Memory Cache in Go – Here’s What I Learned! 🔧🧠

Over the last few days, I’ve been working on a Redis-inspired, lightweight **in-memory caching service written in Go**, aiming to benchmark its performance, implement best concurrency practices, and modularize it for service-level use. Here's a breakdown of my journey:

---

### 🛠️ **What I Built**

A custom **Go-based in-memory key-value store** supporting:

- Basic `Set`, `Get`, `Update`, and `Delete` operations
    
- TTL (Time-To-Live) support with a background cleanup goroutine
    
- Safe concurrency using `sync.RWMutex`
    
- Modular design to support **plug-and-play cache clients**
    

---

### ⚙️ **Key Technical Features**

- 🔁 **TTL Expiry**: Automatic deletion of expired keys every second.
    
- 🔒 **RWMutex over Mutex**: Replaced `sync.Mutex` with `sync.RWMutex` to reduce contention in read-heavy workloads.
    
- ⚡ **Benchmarking**: Used `go test -bench` to validate optimizations.
    
- 📦 **Dependency Injection**: Designed to swap the backend (`inmemcache`, `redis`, etc.) with minimal code changes.
    
- ✅ **Graceful Shutdown**: Catches `SIGINT/SIGTERM` to cleanly stop background goroutines.
    
- 🧪 **Microservice-Ready**: Built as a package usable by any Go service – supports both standalone mode and embedding.
    

---

### 📊 **Performance Benchmark**

Tested on: `Intel i5-8300H`, Go 1.22, Windows/Linux

#### 🔹 `inmemcache` (without RWMutex):

```
BenchmarkSet-8     2221554      515.3 ns/op     222 B/op    4 allocs/op
BenchmarkGet-8     4211239      335.7 ns/op      23 B/op    1 allocs/op
BenchmarkUpdate-8  4299273      319.3 ns/op      23 B/op    1 allocs/op

```

#### 🔸 `rw_memcache` (with RWMutex):

```
BenchmarkSet-8     2216912      520.2 ns/op     222 B/op    4 allocs/op
BenchmarkGet-8     4372797      296.2 ns/op      23 B/op    1 allocs/op ✅
BenchmarkUpdate-8  3985664      328.0 ns/op      23 B/op    1 allocs/op
```

✅ **RWMutex showed ~11% improvement on Get() performance**, ideal for read-heavy use cases!

---

### 🤔 **Lessons Learned**

- RWMutex helps only when you have high read concurrency. Set/Update, being write locks, see little gain.
    
- Pooling or batching operations might help optimize write paths further.
    
- `sync.Map` is good for dynamic workloads but sacrifices type safety.
    
- Graceful shutdown handling is essential even for internal services.
    

---

### 🔄 **Next Steps**

- Add LRU/ARC eviction strategies
    
- Support for persistence (snapshot or AOF-like)
    
- gRPC or HTTP API wrapper for remote usage
    
- Extend with pluggable backends (Redis, Memcached, etc.)
    
