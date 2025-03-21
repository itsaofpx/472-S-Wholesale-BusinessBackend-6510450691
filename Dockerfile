# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# ติดตั้งเครื่องมือที่จำเป็น
RUN apk add --no-cache git

# ตรวจสอบว่ามีไฟล์ go.mod และ go.sum หรือไม่ แล้วค่อยคัดลอก
COPY . .

# สร้างไฟล์ go.mod หากไม่มี
RUN if [ ! -f go.mod ]; then go mod init sa-project; fi

# ดาวน์โหลด dependencies และตรวจสอบ
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:3.18

# Add necessary packages
RUN apk --no-cache add ca-certificates tzdata

# Set timezone to Asia/Bangkok
ENV TZ=Asia/Bangkok

# Create a non-root user to run the application
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy any configuration files if needed
# COPY --from=builder /app/config ./config

# Change ownership of the application files
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
