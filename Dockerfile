FROM golang:1.16.6-alpine3.14 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build -o server server.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/server .

EXPOSE $PORT

CMD ["./server"]
