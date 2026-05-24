# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

<<<<<<< HEAD
RUN go build -o server ./cmd/server/main.go
=======
# Use main's explicit path, or your project's entry point
RUN go build -o server ./main.go
>>>>>>> 5a0dbd76cf1f6923da3ff3d8fa079068a017ac5b

# Runtime stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
