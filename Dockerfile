FROM golang:1.18.2-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/coflnet-bot/main.go


FROM alpine:3.15

COPY --from=builder /app/main /usr/local/bin/coflnet-bot

CMD coflnet-bot