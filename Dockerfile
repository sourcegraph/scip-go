FROM golang:1.20.3-alpine@sha256:08e9c086194875334d606765bd60aa064abd3c215abfbcf5737619110d48d114 as builder

COPY . /sources
WORKDIR /sources
RUN go build -o scip-go ./cmd/scip-go

FROM alpine:latest

COPY --from=builder /sources/scip-go /usr/bin/
CMD ["scip-go"]
