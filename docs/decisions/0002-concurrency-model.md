# ADR-0002: Concurrency Model — Goroutines + Redis Atomicity

## Status
Draft

## Context

The rate limiter must support many concurrent requests safely.

## Decision

Use Go goroutines for concurrency and Redis atomic operations for distributed consistency.

## Alternatives Considered

- OS threads
- Single-threaded processing
- In-memory locking only

## Consequences

- Better scalability
- Easier concurrent request handling
- Requires careful race-condition management