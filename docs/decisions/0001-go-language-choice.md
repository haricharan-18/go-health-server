# ADR-0001: Use Go as the Implementation Language

## Status

Accepted

---

## Context

We needed a language to build a high-performance distributed rate limiter.

The system must handle many concurrent requests efficiently.

---

## Decision

Use Go 1.22.4.

---

## Alternatives Considered

### Python
Easy syntax but limited true concurrency because of GIL.

### Java
Strong ecosystem but heavier runtime and slower startup.

### Node.js
Good for I/O but weaker for CPU-heavy concurrency workloads.

---

## Consequences

- Team must learn Go from scratch
- Goroutines simplify concurrency
- Small binary size
- Fast startup
- Built-in testing and race detector