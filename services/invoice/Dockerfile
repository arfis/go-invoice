FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN ls /go/bin  # This will output the contents of /go/bin to the Docker build logs

RUN go build -o ./tmp/main ./cmd/main/main.go

FROM golang:1.22

WORKDIR /app

COPY --from=builder /app/tmp/main ./tmp/main
COPY --from=builder /go/bin/air /usr/local/bin/air

COPY script/wait-for-it.sh /usr/local/bin/wait-for-it.sh

RUN chmod +x /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/air

CMD ["wait-for-it.sh", "db:5432", "--", "air", "-c", "./config/air.toml"]
