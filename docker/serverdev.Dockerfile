# Build the manager binary
FROM golang:1.22

WORKDIR /workspace

# https://github.com/cosmtrek/air
# used for live reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the Go Modules manifests
COPY server/go.mod      go.mod
COPY server/go.sum      go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY server/cmd         cmd
COPY server/internal    internal
COPY server/modules     modules

# Build
CMD ["air", "--build.cmd", "go build -o bin/api cmd/feature-manager/main.go", "--build.bin", "./bin/api"]
