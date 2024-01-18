# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server .

CMD ["./server"]
