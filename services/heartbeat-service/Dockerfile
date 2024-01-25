FROM golang:1.21-alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /heartbeat-service
CMD ["/heartbeat-service"]
