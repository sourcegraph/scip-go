# When updating the version of the base container, please use the 
# SHA256 of the distribution manifest, not individual images.

# This ensures that when pulling the container, Docker will detect 
# the platform and pull the correct image (if it exists)

# For example, to find out the hash of the manfiest, run:

# $ docker buildx imagetools inspect golang:1.21.5-alpine
# Name:      docker.io/library/golang:1.21.5-alpine
# MediaType: application/vnd.docker.distribution.manifest.list.v2+json
# Digest:    sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b
# And use this digest in FROM
FROM golang:1.21.5@sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b

COPY . /sources
WORKDIR /sources
RUN go build -o scip-go ./cmd/scip-go

FROM alpine:latest@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

COPY --from=builder /sources/scip-go /usr/bin/
CMD ["scip-go"]
