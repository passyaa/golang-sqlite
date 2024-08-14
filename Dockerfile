# Gunakan base image resmi dari Go
FROM golang:1.19-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Install dependencies
RUN go mod tidy

# Build aplikasi Go
RUN go build -o golangApp .

# Ekspos port yang akan digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./golangApp"]