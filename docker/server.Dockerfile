# build the manager binary
FROM golang:1.22 as builder

WORKDIR /workspace

# copy the Go Modules manifests
COPY server/go.mod      go.mod
COPY server/go.sum      go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# copy the go source
COPY server/cmd         cmd
COPY server/internal    internal
COPY server/modules     modules

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager cmd/feature-manager/main.go

# use distroless as minimal base image to package the manager binary
# refer to https://github.com/GoogleContainerTools/distroless for more details
FROM registry.access.redhat.com/ubi8/ubi-minimal

ARG user=1001

WORKDIR /

COPY --from=builder /workspace/manager .

USER ${user}

ENTRYPOINT ["/manager"]
