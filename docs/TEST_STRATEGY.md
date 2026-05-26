# Test Strategy — sei-ratelimiter

## Test Types

### Unit Tests
Location: internal/algorithms/*_test.go

Requirement:
- No external dependencies
- Run offline

Command:
go test -race -v ./internal/algorithms/...

---

### Integration Tests
Location:
- internal/store/*_test.go
- integration/...

Requirement:
- Redis must be running

Command:
REDIS_URL=localhost:6379 go test -race -v ./...

Behaviour when Redis missing:
- Tests SKIP
- Tests do NOT FAIL

---

### Load Tests
Location: tests/load/*.js

Tool: k6

Requirement:
- Full docker compose stack running

Command:
k6 run tests/load/smoke.js

---

## Redis for Tests

We use a manually managed Redis container instead of testcontainers-go.

Reason:
- Simpler setup
- Faster execution
- Less dependency complexity

---

## Required Flags

Every Go test command must include:

-race

A test without -race is not considered complete.