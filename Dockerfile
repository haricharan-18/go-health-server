FROM golang:1.26-alpine

WORKDIR /app

COPY . .

RUN go build -o health-server ./cmd/server

EXPOSE 8080

CMD ["./health-server"]
