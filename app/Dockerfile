FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

# Expose the application on port 8080
EXPOSE 8080

# Run the application
CMD ["./app"]
