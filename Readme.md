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
Set:      515.3 ns/op
Get:      296.2 ns/op
Update:   319.3 ns/op
```

#### 🔸 `rw_memcache` (with RWMutex):

```
Set:      520.2 ns/op
Get:      296.2 ns/op ✅
Update:   328.0 ns/op
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
    

---

### 📂 Repo (Coming Soon)

Publishing this soon on GitHub with clean structure and benchmark results. Want early access or want to contribute? Let me know! 👇

---

If you're exploring **system design**, **Go microservices**, or **high-performance in-memory caching**, I’d love to connect and chat more.

#golang #redis #microservices #opensource #performance #programming #systemdesign #backend #concurrency #devprojects