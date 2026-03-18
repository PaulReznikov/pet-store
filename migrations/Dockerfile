FROM golang:1.26.1-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY migrations/ ./migrations/

# CGO_ENABLED=0 делает бинарник статически слинкованным, чтобы он работал в пустом alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrator-app ./migrations/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /migrator-app .

CMD ["./migrator-app"]