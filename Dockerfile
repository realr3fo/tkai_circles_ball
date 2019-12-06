# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container (GOPATH/src/app_name)
WORKDIR $GOPATH/src/github.com/realr3fo/tkai_circles_ball

# Copy the local package files to the container's workspace.
ADD . $GOPATH/src/github.com/realr3fo/tkai_circles_ball

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy Gopkg.toml Gopkg.lock files
COPY Gopkg.toml Gopkg.lock ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN dep ensure --vendor-only

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/realr3fo/tkai_circles_ball

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/tkai_circles_ball

# Expose port 8080 to the outside world
EXPOSE 8002
