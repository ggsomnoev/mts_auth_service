FROM golang:1.24 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
COPY vendor ./vendor
ENV GOFLAGS=-mod=vendor

COPY internal ./internal
COPY cmd/cnvalidator ./cmd/cnvalidator

COPY certs ./certs

RUN go build -o cnvalidator ./cmd/cnvalidator

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/cnvalidator .
COPY --from=builder /build/certs ./certs

CMD ["./cnvalidator"]