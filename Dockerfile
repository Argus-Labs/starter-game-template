################################
# Build Image - Normal
################################
FROM golang:1.22-bookworm AS build

WORKDIR /go/src/app

# Copy the go module files and download the dependencies
# We do this before copying the rest of the source code to avoid
# having to re-download the dependencies every time we build the image
COPY /cardinal/go.mod /cardinal/go.sum ./
RUN go mod download

# Set the GOCACHE environment variable to /root/.cache/go-build to speed up build
ENV GOCACHE=/root/.cache/go-build

# Copy the rest of the source code and build the binary
COPY /cardinal ./
RUN --mount=type=cache,target="/root/.cache/go-build" go build -v -o /go/bin/app

################################
# Runtime Image - Normal
################################
FROM gcr.io/distroless/base-debian12 AS runtime

# Copy world.toml to the image
COPY world.toml world.toml

# Copy the binary from the build image
COPY --from=build /go/bin/app /usr/bin

# Run the binary
CMD ["app"]

################################
# Runtime Image - Debug
################################
FROM golang:1.22-bookworm AS runtime-debug

WORKDIR /go/src/app

# Install delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Copy the go module files and download the dependencies
# We do this before copying the rest of the source code to avoid
# having to re-download the dependencies every time we build the image
COPY /cardinal/go.mod /cardinal/go.sum ./
RUN go mod download

# Set the GOCACHE environment variable to /root/.cache/go-build to speed up build
ENV GOCACHE=/root/.cache/go-build

# Copy the rest of the source code and build the binary with debugging symbols
COPY /cardinal ./
RUN --mount=type=cache,target="/root/.cache/go-build" go build -gcflags="all=-N -l" -v -o /usr/bin/app

# Copy world.toml to the image
COPY world.toml world.toml

CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/usr/bin/app"]