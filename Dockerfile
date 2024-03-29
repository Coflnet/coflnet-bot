FROM registry.suse.com/bci/golang:1.22 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM registry.suse.com/bci/bci-micro:15.5

COPY --from=builder /build/bin/app /app

ENTRYPOINT ["/app"]
