# Dockerfile for Go service
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main ./cmd/api
CMD ["./main"]
