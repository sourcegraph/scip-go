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
#
# If you have the skopeo CLI, you can use:
#
# $ skopeo inspect --raw docker://golang:1.25.0 | sha256sum

ARG base_sha=5502b0e56fca23feba76dbc5387ba59c593c02ccc2f0f7355871ea9a0852cebe

FROM golang:1.25.0@sha256:${base_sha} AS builder

COPY . /sources
WORKDIR /sources
RUN go build -o scip-go ./cmd/scip-go

# Keep in sync with builder image
FROM golang:1.25.0@sha256:${base_sha} AS final

COPY --from=builder /sources/scip-go /usr/bin/
ENV GOTOOLCHAIN=auto
CMD ["scip-go"]
