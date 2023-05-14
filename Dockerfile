FROM golang:1.18-alpine AS builder

WORKDIR /tg-bot

COPY . .

RUN go build -v -o ./bin/tg-bot ./cmd/main.go

ARG token
ENV TELEGRAM_TOKEN=$token

ENTRYPOINT ["./bin/tg-bot"]