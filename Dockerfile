# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /student-list-app

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /app
COPY --from=builder /student-list-app /app/student-list-app

ENV PORT=8080

EXPOSE 8080

CMD ["/app/student-list-app"]