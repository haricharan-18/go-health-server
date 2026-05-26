# ADR-0002: Infrastructure Tooling — Docker Compose + Redis 7 + Multi-Stage Builds

## Status

Accepted

## Context

All 5 team members need a consistent local development environment.

The rate limiter requires two app instances sharing one Redis.

Setup must work with a single command on any team member's machine.

## Decision

- Docker Compose v2 for local orchestration
- Redis 7-alpine as the shared datastore
- Multi-stage Dockerfile (golang:1.22-alpine builder + alpine:3.19 runtime)
- redis_data volume for data persistence across restarts

## Alternatives Considered

- Manual docker run commands: too much setup and error-prone
- Kubernetes (minikube): too complex for this project stage
- Redis non-alpine: larger image size with no extra benefit
- Single-stage Dockerfile: creates much larger images

## Consequences

Good:
- One-command startup using docker compose up --build
- Smaller image size with multi-stage builds
- Persistent Redis storage using volumes

Bad:
- Redis is a single point of failure in local setup

Accepted for now.