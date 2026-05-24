# ADR-0003: Package Structure

## Status

Accepted

## Context

The rate limiter needs clear separation between HTTP handling,
rate limiting logic, and Redis operations.

Without this separation, testing algorithms requires a live HTTP server
and a real Redis connection — making tests slow and fragile.

---

## Decision

Use a layered package structure inside internal/:

internal/algorithms — pure logic, depends on Store interface only

internal/store — Redis implementation of Store interface

internal/api — HTTP handlers only

internal/config — environment variables only

cmd/server — composition root, no logic

---

## Alternatives Considered

Flat package (everything in main):

Rejected. Untestable — impossible to mock Redis or the HTTP layer.

Feature-based packages (ratelimit/, rules/):

Rejected. Creates circular dependencies when features share types.

Hexagonal architecture (ports and adapters):

Considered but over-engineered for this scale.

Our layered approach achieves the same testability goal more simply.

---

## Consequences

Good:

Algorithms testable without Redis (inject a fake Store).

HTTP handlers testable without real algorithms (inject a fake Limiter).

Clear dependency direction enforced by the compiler.

No circular imports possible with this layout.

Bad:

More files and directories upfront.

Team must understand dependency injection.