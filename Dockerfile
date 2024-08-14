# Gunakan base image resmi dari Go dengan CGO diaktifkan
FROM golang:1.23-alpine

# Instal dependensi untuk CGO (misalnya, build-essential untuk compilers)
RUN apk add --no-cache gcc musl-dev

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Aktifkan CGO dan set target untuk build
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

# Install dependencies
RUN go mod tidy

# Build aplikasi Go
RUN go build -o golangApp .

# Ekspos port yang akan digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./golangApp"]
