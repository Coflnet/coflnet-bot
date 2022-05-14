FROM golang:1.18.2-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD /app/main