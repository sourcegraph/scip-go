# When updating the version of the base container, please use the 
# SHA256 listed under 'Index digest' on Docker Hub,
# not the 'Manifest digest'.
#
# This ensures that when pulling the container, Docker will detect 
# the platform and pull the correct image (if it exists)
#
# Alternate way of determining the Index digest using the docker CLI.
#
# $ docker buildx imagetools inspect golang:1.21.5-alpine
# Name:      docker.io/library/golang:1.21.5-alpine
# MediaType: application/vnd.docker.distribution.manifest.list.v2+json
# Digest:    sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b
# Manifests:
#    <other stuff to ignore>
# And use this digest in FROM

ARG base_sha=adee809c2d0009a4199a11a1b2618990b244c6515149fe609e2788ddf164bd10

FROM golang:1.23.2@sha256:${base_sha} as builder

COPY . /sources
WORKDIR /sources
RUN go build -o scip-go ./cmd/scip-go

# Keep in sync with builder image
FROM golang:1.23.2@sha256:${base_sha} as final

COPY --from=builder /sources/scip-go /usr/bin/
CMD ["scip-go"]
