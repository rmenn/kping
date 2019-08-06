FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/kping /app/

ENTRYPOINT ./kping
