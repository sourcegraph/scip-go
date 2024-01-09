FROM golang:1.21.5@sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b

COPY . /sources
WORKDIR /sources
RUN go build -o scip-go ./cmd/scip-go

FROM alpine:latest

COPY --from=builder /sources/scip-go /usr/bin/
CMD ["scip-go"]
