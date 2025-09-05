# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make gcc musl-dev

# Копируем go.mod и go.sum
COPY back/go.mod back/go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY back/. .

# Сборка бинарника для Linux
RUN GOOS=linux GOARCH=amd64 go build -o back ./cmd/server/main.go && ls -l /app

# Stage 2: Run
FROM alpine:3.18

WORKDIR /app

# Устанавливаем бинарник в безопасное место
COPY --from=builder /app/back /usr/local/bin/back

ENV PORT=3000
EXPOSE 3000

CMD ["/usr/local/bin/back"]
