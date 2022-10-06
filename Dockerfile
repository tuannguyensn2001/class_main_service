FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go build -o main src/server/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 80
CMD ["/app/main","server"]