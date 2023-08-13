FROM golang:1.21-bullseye as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p bin && go build -o ./bin ./...

FROM gcr.io/distroless/base-debian11

COPY --from=builder /build/bin/coflnet-bot /app

ENTRYPOINT ["/app"]
