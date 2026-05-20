# Shared State in the Rate Limiter

## In-Memory Version

Shared state:
fw.counts map[string]int

Protected by:
sync.Mutex in each FixedWindow struct

Risk:
Works correctly only on a single node.
Multiple nodes cannot share in-memory state.

---

## Redis-Backed Version

Shared state:
Redis key per clientID

Protected by:
Redis atomic operations

Risk:
INCR and EXPIRE are separate operations.
Two nodes may update simultaneously.

---

## Lua Script Version

Shared state:
Redis key per clientID

Protected by:
Atomic Lua script execution

Why it works:
Redis executes Lua scripts atomically without interruption.