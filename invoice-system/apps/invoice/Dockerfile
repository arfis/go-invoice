FROM golang:1.22 as builder

WORKDIR /app

# Copy necessary files
COPY ./apps/invoice /app
COPY ./commands /commands

# Modify go.mod for Docker environment paths
RUN sed -i 's|example.com/commands => ../../commands|example.com/commands => /commands|g' go.mod

# Install dependencies
RUN go mod download

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@latest
RUN ls /go/bin  # List contents to verify air installation

# Build your Go application
RUN go build -o /app/main ./cmd/main/main.go

FROM golang:1.22

WORKDIR /app

# Copy the compiled main application
COPY --from=builder /app/main /app/main

# Try copying air from the builder stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Other setup...
COPY ./apps/invoice/script/wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/air

CMD ["wait-for-it.sh", "db:5432", "--", "air", "-c", "./config/air.toml"]
