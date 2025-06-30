FROM golang:1.24.4 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main /app/main

RUN chmod +x /app/main

CMD ["/app/main"]
