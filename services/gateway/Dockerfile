FROM golang:1.22

WORKDIR /app

# Explicitly set PATH to ensure it's correctly recognized
ENV PATH="/go/bin:${PATH}"

# Install necessary tools
RUN go install github.com/cosmtrek/air@latest

# Copy everything from the current directory to /invoice in the container
COPY . .

# Download all dependencies
RUN go mod download

# Build your application
RUN go build -o ./tmp/main ./cmd/main/main.go

# Set air to run with the specified configuration file
CMD ["/go/bin/air", "-c", "./config/air.toml"]
