# Docker Concepts — sei-ratelimiter

## Image

A Docker image is a blueprint used to create containers. It contains the application code, dependencies, runtime, and operating system libraries required to run the application. Images are built using a Dockerfile and remain immutable after creation. Multiple containers can be created from the same image.

## Container

A container is a running instance of a Docker image. It provides an isolated environment where the application runs independently from the host machine. Containers have their own filesystem, processes, and networking. Multiple containers from the same image can run simultaneously.

## Layer and Caching

Each instruction inside a Dockerfile creates a layer. Docker caches these layers to improve build performance. If a layer does not change, Docker reuses the cached version during rebuilds. This is why dependency files like go.mod are copied before source code to avoid downloading dependencies repeatedly.

## Volume

A Docker volume provides persistent storage outside the container lifecycle. Even if a container is deleted or restarted, the volume data remains safe. In this project, Redis uses a volume to persist counters and rules across restarts.

## Network

Docker Compose automatically creates a shared network for all services. Containers inside this network communicate using service names instead of IP addresses. This allows app containers to connect to Redis using REDIS_URL=redis:6379.