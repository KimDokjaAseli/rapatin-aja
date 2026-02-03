# Stage 1: Builder
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy SEMUA file (menggunakan . bukan rapatk/)
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Stage 2: Final image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN adduser -D appuser
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x main
USER appuser

EXPOSE 8080
CMD ["./main"]