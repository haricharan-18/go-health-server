# Sprint Log — Zartex SEI Project 1

---

## Day 1 — May 18, 2026

**Phase:** 1 — Foundation
**Goal:** Go fundamentals + HTTP /health server
**Schedule:** All 5 members, same tasks, no role split today

### Completed

- All 5 members completed tour.golang.org sections 1–13
  (variables, types, functions, structs, pointers, error handling, methods)
- All 5 members read Go by Example: error handling, closures, interfaces
- All 5 members built HTTP server: GET /health returns {"status": "ok"}
- All 5 members tested server with curl — all returning correct JSON
- EOD confidence ratings posted in group chat

### What Was Harder Than Expected

- Pointers — the difference between value and address took time to click
- Understanding why pointer receivers are needed on struct methods
- Interfaces — the concept of implicit implementation (no 'implements' keyword)

### What Was Easier Than Expected

- Setting up the HTTP server — Go's net/http package is straightforward
- JSON encoding — json.NewEncoder(w).Encode() just works

### Phase 1 Gate Status

- [ ] All 5 members: HTTP server returns {"status":"ok"} on GET /health
- [ ] All 5 members: Go fundamentals tour complete

---

## Day 2 — May 19, 2026

**Phase:** 1 — Foundation
**Goal:** Full dev environment + GitHub repo + role-specific tasks
**Schedule:** All 5 members

### Environment Setup (All Members)

- [ ] WSL2 Ubuntu 22.04 installed and running
- [ ] Go 1.22.4 installed and verified (go version)
- [ ] VS Code connected to WSL2 (WSL: Ubuntu in bottom-left)
- [ ] Git configured (name, email, defaultBranch=main)
- [ ] GitHub SSH key generated and added to account
- [ ] sei-ratelimiter repo cloned successfully

### Role Tasks

- [ ] Abhishek: k6 installed, healthHandler refactored, 4 tests passing, PR submitted
- [ ] Madhu: Limiter interface, Config struct, FixedWindow impl, PR submitted
- [ ] Gayathri: 5 handler tests written and passing, PR submitted
- [ ] Hari: Docker Desktop installed, Redis CLI commands practiced, PR submitted
- [ ] Vishnu: README skeleton, ADR template, ADR-001, Sprint Log, PR submitted

### PRs Submitted Today

| Member | Branch | Status |
|--------|--------|--------|
| Abhishek | day2/abhishek-setup | |
| Madhu | day2/madhu-interface | |
| Gayathri | day2/gayathri-tests | |
| Hari | day2/hari-docker-redis | |
| Vishnu | day2/vishnu-docs | |

### Notes

<!-- Fill in after EOD -->
<!-- What was harder than expected? -->
<!-- Did everyone get environment set up? Any blockers? -->

---