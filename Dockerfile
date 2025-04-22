# Dockerfile for Go service
FROM golang:1.23.0
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .
CMD ["./main"]
