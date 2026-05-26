# Test Conventions — sei-ratelimiter
 
## Test Function Naming
Pattern: TestFunctionName_Condition
Examples:
  TestFixedWindow_AllowsUnderLimit
  TestFixedWindow_BlocksAtLimit
  TestCheckHandler_Returns400ForMissingClientID
 
## Required Flags
Every go test run includes -race. No exceptions.
  make test          → go test -race -v ./...
  make test-count    → go test -race -count=3 ./...
 
## Test Types
 
Unit tests:
  No external dependencies. Run offline.
  Location: internal/algorithms/*_test.go
  Package: package algorithms_test (black-box)
 
Integration tests:
  Require Redis. Skip gracefully if Redis is not running.
  Use testhelpers.RedisClient(t) — never create a raw redis.Client in tests.
  Always clean up keys in t.Cleanup() using testhelpers.FlushKeys().
 
Load tests:
  Location: tests/load/*.js
  Tool: k6
  Require full docker compose stack.
 
## Table-Driven Pattern
Use for any test with more than one case.
 
  tests := []struct {
      name     string
      input    SomeType
      expected SomeType
  }{
      {"case description", input, expected},
  }
  for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
          // test body
      })
  }
 
## t.Fatal vs t.Error
t.Fatal  — stops the test immediately. Use when continuing makes no sense.
t.Error  — marks failed but continues. Use when multiple independent checks exist.
t.Skip   — skips test. Use when prerequisite (Redis) is not available.
