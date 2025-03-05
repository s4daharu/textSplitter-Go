# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o text-splitter .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/text-splitter .
COPY templates/ ./templates/
EXPOSE 8080
CMD ["./text-splitter"]