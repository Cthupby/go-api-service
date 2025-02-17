# Build stage
FROM golang:alpine AS build-stage

# Set working directory
WORKDIR /app

# Copy depencies and app directories
ADD go.mod .
COPY . .

# Build application
RUN go build -o apiservice  cmd/apiservice/main.go

# Production stage
FROM golang:alpine AS production-stage

# Set working directory
WORKDIR /app

# Copy built application
COPY --from=build-stage /app/apiservice /app/apiservice

# Expose port and run application
EXPOSE 5000
CMD ["./apiservice"]
