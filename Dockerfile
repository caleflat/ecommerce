FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/bin/ ./cmd/...

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/bin/ /app/bin/

ENTRYPOINT ["/app/bin/"]

ARG APP

CMD ["server", "${APP}"]
