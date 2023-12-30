FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/safblog-backend ./cmd/main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/safblog-backend .
EXPOSE 8080
CMD ["./safblog-backend"]
