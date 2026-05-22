# Go Concurrency

## Goroutines vs Threads

Goroutines are lightweight concurrent functions managed by the Go runtime.

## Race Conditions

A race condition happens when multiple goroutines modify shared data at the same time.

## sync.Mutex

Mutex protects shared data by allowing only one goroutine to access critical code at a time.

## Why Redis Still Needs Lua Scripts

Mutex works only inside one Go process.
Distributed systems need Redis atomic operations to avoid cross-server race conditions.