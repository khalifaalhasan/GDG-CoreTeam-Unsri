# --- STAGE 1: BUILDER ---
# Menggunakan image resmi Go untuk proses kompilasi
FROM golang:1.21-alpine AS builder

# Set direktori kerja di dalam container
WORKDIR /app

# Copy go.mod dan go.sum untuk mengunduh dependencies (efisien caching)
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy seluruh kode sumber ke container
COPY . .

# Kompilasi aplikasi Go
# -o /app/server : output binary bernama 'server' di root /app
# -ldflags -s -w : mengurangi ukuran binary
# ./cmd/api/main.go : lokasi file utama Anda
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api/main.go


# --- STAGE 2: FINAL IMAGE (Minimal dan Aman) ---
# Menggunakan image 'scratch' (image paling dasar yang hanya berisi binary)
FROM scratch

# Set variabel environment (untuk menjalankan server)
ENV PORT 8080 
# Cloud Run menggunakan PORT 8080 secara default, tetapi ini adalah praktik terbaik.

# Salin binary yang sudah dikompilasi dari stage 'builder'
COPY --from=builder /app/server /server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Command untuk menjalankan aplikasi (binary 'server')
ENTRYPOINT ["/server"]