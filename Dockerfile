# Build Stage
FROM golang:1.24.0-alpine AS builder
WORKDIR /app

# Install bash and any other dependencies
RUN apk add --no-cache bash

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and wait-for-it script
COPY . ./
COPY wait-for-it.sh /app/wait-for-it.sh  
# Ensure the script is copied into the container

RUN go build -o main .

# Runtime Stage
FROM alpine:latest
WORKDIR /root/

# Install bash (since we need bash in the runtime container)
RUN apk add --no-cache bash

# Copy the backend binary and .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

# Copy the wait-for-it script from the builder stage
COPY --from=builder /app/wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh 
# Ensure the script is executable

EXPOSE ${APP_PORT}

# Run wait-for-it before starting the backend
CMD ["/app/wait-for-it.sh", "db:5432", "--", "./main"]