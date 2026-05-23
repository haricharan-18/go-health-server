# sei-ratelimiter

Distributed Rate Limiter as a Service — Zartex SEI Project 1

---

## Architecture

![Architecture](docs/architecture-v1.png)

---

## Algorithms

- Fixed Window
- Sliding Window
- Token Bucket

---

## API Reference

### POST /check

Checks whether request is allowed.

### POST /rules

Create a new rule.

### GET /rules

Get all rules.

### GET /rules/:id

Get rule by ID.

### DELETE /rules/:id

Delete rule.

---

## How To Run

Coming soon.

---

## How To Run Tests

Coming soon.

---

## Benchmarks

Coming soon.

---

## Failure Modes

Coming soon.

---

## What We Would Do at 10x Scale

Coming soon.


## How To Run

### Prerequisites

- Docker Desktop
- WSL2 enabled
- Git

### Start the full stack

```bash
git clone git@github.com:Zartex-the-art/sei-ratelimiter.git
cd sei-ratelimiter
docker compose up --build
```

This starts:

- App Node 1 → http://localhost:8080
- App Node 2 → http://localhost:8081
- Redis → localhost:6379

### Verify

```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
```

Both should return:

```json
{"status":"ok"}
```

### Stop

```bash
docker compose down
```

## Project Structure

sei-ratelimiter/
├── cmd/
│
├── docs/
│   ├── decisions/
│   │   ├── 0000-template.md
│   │   ├── 0001-go-language-choice.md
│   │   ├── 0002-concurrency-model.md
│   │   ├── 0002-infrastructure-tooling.md
│   │   └── 0003-package-structure.md
│   │
│   ├── diagrams/
│   │   ├── architecture-v1.png
│   │   ├── architecture-v2.png
│   │   └── architecture-v3.png
│   │
│   ├── ARCHITECTURE.md
│   ├── CONCURRENCY.md
│   ├── DOCKER_CONCEPTS.md
│   ├── redis_race_notes.md
│   └── SHARED_STATE.md
│
├── internal/
│   └── algorithms/
│       ├── fixed_window.go
│       ├── fixed_window_test.go
│       └── limiter.go
│
├── pkg/
│
├── practice/
│   ├── goroutine.go
│   ├── mutex.go
│   └── race.go
│
├── go.mod
├── main.go
├── README.md
├── server_test.go
└── SPRINT_LOG.md