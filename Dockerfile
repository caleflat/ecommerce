FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

CMD ["/app/main"]