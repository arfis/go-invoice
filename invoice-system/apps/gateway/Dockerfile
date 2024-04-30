FROM golang:1.22

WORKDIR /app

# Copy necessary files
COPY ./apps/gateway /app
COPY ./commands /commands

# Modify go.mod for Docker environment paths
RUN sed -i 's|example.com/commands => ../../commands|example.com/commands => /commands|g' go.mod


# Explicitly set PATH to ensure it's correctly recognized
ENV PATH="/go/bin:${PATH}"

# Install necessary tools
RUN go install github.com/cosmtrek/air@latest


# Download all dependencies
RUN go mod download

# Build your application
RUN go build -o ./tmp/main ./cmd/main/main.go

# Set air to run with the specified configuration file
CMD ["/go/bin/air", "-c", "./config/air.toml"]
