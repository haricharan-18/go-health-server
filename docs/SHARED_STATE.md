## In-Memory Version

Shared state:
`fw.counts` → `map[string]int` (per client request count)

Protected by:
`sync.Mutex` inside the `FixedWindow` struct

How it works:
- All reads and writes to the map are protected by a mutex
- Only one goroutine can access the map at a time

Risk:
Works correctly only on a single node.
Multiple nodes cannot share in-memory state.

---

## Redis-Backed Version

Shared state:
Redis key per client (e.g., `rate_limit:&lt;clientID&gt;`)

Protected by:
Redis commands like `INCR`

Problem:
`INCR` and `EXPIRE` are separate operations.
Between these operations, another request can run.

Race condition example:
- Node A reads value → 99
- Node B reads value → 99
- Node A sets → 100
- Node B sets → 100 (should be 101)

Result:
Lost updates. Incorrect rate limiting (allows more requests than expected).

---

## Lua Script Version

Shared state:
Redis key per client

Protected by:
Atomic Lua script execution

Why it works:
Redis executes Lua scripts atomically without interruption.

What becomes atomic:
- Read (`GET`)
- Modify (`INCR` / logic)
- Expire (`SET` TTL)

Result:
No race conditions. Correct behavior across multiple nodes.

---

## Key Takeaway

- `sync.Mutex` solves concurrency inside one node
- Redis alone does NOT guarantee full correctness
- Lua scripts are required for atomic operations across distributed systems