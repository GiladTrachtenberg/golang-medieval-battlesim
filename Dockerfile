FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
COPY go.mod .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
