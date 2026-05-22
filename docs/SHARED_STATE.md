# Shared State in the Rate Limiter

## In-Memory Version (Current — Early Phase)

**Shared State:**
- `fw.counts` → map[string]int (per client request count)

**Protection Mechanism:**
- `sync.Mutex` inside the FixedWindow struct

**How it works:**
- All reads and writes to the map are protected by a mutex
- Only one goroutine can access the map at a time

**Limitation:**
- Works correctly only on a single node
- If multiple app instances run, each has its own separate memory
- No shared state across nodes → inconsistent rate limiting


---

## Redis-Backed Version (Next Phase)

**Shared State:**
- Redis key per client (e.g., `rate_limit:<clientID>`)

**Protection Mechanism:**
- Redis commands like `INCR`

**Problem:**
- Operations like `INCR` and `EXPIRE` are separate
- Between these operations, another request can run

**Race Condition Example:**
- Node A reads value → 99
- Node B reads value → 99
- Node A sets → 100
- Node B sets → 100 (should be 101)

**Result:**
- Lost updates
- Incorrect rate limiting (allows more requests than expected)


---

## Lua Script Version (Final — Correct Solution)

**Shared State:**
- Redis key per client

**Protection Mechanism:**
- Lua script executed atomically in Redis

**Why it works:**
- Redis executes Lua scripts as a single operation
- No other command runs in between

**What becomes atomic:**
- Read (GET)
- Modify (INCR / logic)
- Expire (SET TTL)

**Result:**
- No race conditions
- Correct behavior across multiple nodes


---

## Key Takeaway

- `sync.Mutex` solves concurrency inside one node
- Redis alone does NOT guarantee full correctness
- Lua scripts are required for atomic operations across distributed systems