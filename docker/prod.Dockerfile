FROM golang:1.22 AS builder
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum
COPY go.mod go.sum ./
RUN go mod download

# คัดลอกโค้ดแอปพลิเคชัน
COPY . .

# สร้างแอปพลิเคชัน
RUN go build -o myapp ./cmd/main.go

# ขั้นตอนการสร้างภาพสุดท้าย
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
