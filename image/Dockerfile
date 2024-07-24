# Server
FROM registry.access.redhat.com/ubi9/go-toolset:1.21.11-7 as server

WORKDIR /go/src/github.com/stackrox/release-registry

COPY cmd cmd
COPY pkg pkg
COPY gen gen
COPY go.mod .
COPY go.sum .
COPY Makefile Makefile

RUN go build -o build/release-registry cmd/server/main.go

# # UI
# FROM node:18.15.0 as ui

# COPY ui ui

# WORKDIR /ui

# RUN --mount=type=secret,id=npmrc,target=/root/.npmrc make

# Application
FROM registry.access.redhat.com/ubi8/ubi-minimal:8.8-1014 as app

# COPY --from=server /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=server /go/src/github.com/stackrox/release-registry/build/release-registry /release-registry

COPY --from=server /go/src/github.com/stackrox/release-registry/gen/openapiv2 docs

# # COPY --from=ui ui/build ui

ENTRYPOINT [ "/release-registry" ]
