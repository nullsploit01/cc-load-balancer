FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY server .

RUN go build -o http-server .

EXPOSE 8080

CMD ["./http-server"]