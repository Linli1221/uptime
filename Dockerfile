# --- Frontend Build Stage ---
FROM node:20-alpine AS frontend-builder

WORKDIR /app

# Copy frontend code
COPY frontend/ .

# Install dependencies
RUN npm install

# Build the Next.js application for static export
RUN npm run build

# --- Backend Build Stage ---
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

# Install build dependencies
RUN apk --no-cache add build-base

# Copy backend go module files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the rest of the backend source code
COPY backend/ ./

# Copy the built frontend static files from the frontend-builder stage
COPY --from=frontend-builder /app/frontend/out ./frontend_dist

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# --- Final Runtime Stage ---
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates sqlite

WORKDIR /root/

# Copy the built Go binary from the backend-builder stage
COPY --from=backend-builder /app/main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]