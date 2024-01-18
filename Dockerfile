# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server .

CMD ["./server"]
