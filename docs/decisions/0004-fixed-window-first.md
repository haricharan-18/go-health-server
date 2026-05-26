# ADR-004: Fixed Window First

## Status
Accepted

## Context
The project requires a simple and understandable rate limiting algorithm before implementing more advanced algorithms.

## Decision
We will implement the Fixed Window algorithm first.

## Rationale
- Simplest rate limiting algorithm
- Easy to test and debug
- Establishes the Limiter interface contract
- Test patterns can be reused later
- Good teaching foundation for distributed rate limiting

## Known Tradeoff
Fixed Window suffers from the boundary burst problem.

## Future Work
Phase 4 will improve atomicity using Lua scripts.