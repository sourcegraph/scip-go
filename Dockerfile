FROM golang:1.20.3-alpine as builder

COPY . /sources
WORKDIR /sources
RUN go build -o scip-go ./cmd/scip-go

FROM alpine:latest

COPY --from=builder /sources/scip-go /usr/bin/
CMD ["scip-go"]
