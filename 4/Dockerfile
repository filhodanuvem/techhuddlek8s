# Stage 1: Build the Go application
FROM golang:alpine3.19 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Create the final lightweight image
FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
