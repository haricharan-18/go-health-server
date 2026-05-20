# Redis Race Condition Notes

## Overview
Redis commands are atomic individually, but **multi-step operations** (read → compute → write) are NOT atomic. In concurrent environments, this creates race conditions.

---

## Common Redis Race Patterns

### 1. Read-Modify-Write
```
GET counter          → 100
# Another client increments here
SET counter 101      → Should be 102, but it's 101!
```

**Fix:** Use atomic commands
```
INCR counter         → Always correct
```

### 2. Check-and-Set (e.g., stock inventory)
```
GET stock:item123    → 5
# Another client buys 3
DECRBY stock:item123 2  → Stock is now -1 (oversold!)
```

**Fix:** Use Lua script for atomic check-and-act
```lua
local stock = redis.call('GET', KEYS[1])
if tonumber(stock) >= tonumber(ARGV[1]) then
    redis.call('DECRBY', KEYS[1], ARGV[1])
    return 1
else
    return 0
end
```

### 3. Rate Limiting
```
LLEN requests:user1  → 99
# 10 requests arrive simultaneously
LPUSH requests:user1 ...
LTRIM requests:user1 0 99
```

**Fix:** Use Redis Sorted Sets with `ZADD` + `ZREMRANGEBYSCORE`

---

## Go + Redis Best Practices

### Use Pipeline for Batches
```go
pipe := rdb.Pipeline()
pipe.Incr(ctx, "counter")
pipe.Expire(ctx, "counter", time.Hour)
_, err := pipe.Exec(ctx)
```

### Use WATCH for Optimistic Locking
```go
err := rdb.Watch(ctx, func(tx *redis.Tx) error {
    n, _ := tx.Get(ctx, "key").Int()
    _, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
        pipe.Set(ctx, "key", n+1, 0)
        return nil
    })
    return err
}, "key")
```

### Use Distributed Locks for Critical Sections
```go
lock := redislock.New(rdb)
mutex, err := lock.Obtain(ctx, "my-lock", 10*time.Second, nil)
if err != nil {
    return err
}
defer mutex.Release(ctx)
// Critical section here
```

---

## Testing for Redis Races

### Concurrent Client Simulation
```go
func TestRedisIncr(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            rdb.Incr(ctx, "counter")
        }()
    }
    wg.Wait()
    // counter must be exactly 100
}
```

### Load Testing
```bash
redis-benchmark -t set,get -n 100000 -c 50 -P 16
```

---

## Key Takeaways

| Pattern | Risk | Solution |
|---------|------|----------|
| GET → SET | High | `INCR`, `DECR`, `APPEND` |
| Check then Act | High | Lua scripts or `WATCH` |
| Multiple writes | Medium | `Pipeline` or `MULTI/EXEC` |
| Critical sections | High | Distributed locks |

> **Golden Rule:** If you read a value before writing, the operation is NOT atomic. Use Lua, transactions, or single atomic commands.