# Architecture — sei-ratelimiter

## Overview

sei-ratelimiter is a distributed rate limiter exposed as an HTTP service.

Two identical app nodes share one Redis instance.

All rate limit state lives in Redis — no node-local state.

Either node can serve any client's requests correctly.

---

## System Diagram

Client → :8080 → App Node 1 ─┐
                             ├─→ Redis :6379
Client → :8081 → App Node 2 ─┘

---

## Request Flow

POST /check
↓
internal/api — decode request, validate fields
↓
Config resolution — does this clientID have a stored rule?
↓
internal/algorithms — NewLimiter(algorithm, store, limit, windowSecs)
↓
Limiter.Allow(ctx, clientID) — reads/writes Redis via Store interface
↓
Response: {allowed: bool, remaining: int}

---

## Package Layers

cmd/server

Entry point. Reads config. Wires packages together.

No business logic.

internal/api

HTTP layer. Knows HTTP. Calls algorithms.

No direct Redis. No algorithm logic.

internal/algorithms

Rate limiting logic. Knows nothing about HTTP.

Depends on Store interface — not on Redis directly.

Contains: FixedWindow, SlidingWindow, TokenBucket.

internal/store

Redis implementation of Store interface.

Only package that imports go-redis.

Algorithms receive Store via dependency injection.

internal/config

Environment variable loading. No logic.

---

## Dependency Direction

cmd/server → internal/api → internal/algorithms → internal/store → Redis

cmd/server → internal/config

No circular imports possible.

---

## Algorithms

Fixed Window — INCR + EXPIRE. Simple. Boundary burst problem.

Sliding Window — ZADD + ZREMRANGEBYSCORE + ZCOUNT. No boundary burst.

Token Bucket — HSET + HGETALL. Allows burst up to limit.

---

## Distributed Correctness

All Redis operations in Phase 4 become Lua scripts.

Redis executes Lua atomically — no interleaving between nodes.

Prevents over-counting when two nodes serve the same clientID simultaneously.

---

## Failure Modes

Redis unavailable → /check returns HTTP 503 (not 429).

App node restart → other node continues serving, shared Redis state intact.