# Build stage
FROM golang:alpine3.20 AS builder
WORKDIR /app

# Menyalin go.mod dan go.sum dan mengunduh dependensi
COPY go.mod go.sum ./
RUN go mod download

# Menyalin sumber kode aplikasi ke dalam image
COPY . .

# Membangun aplikasi Go; pastikan untuk menyesuaikan path jika struktur folder berbeda
RUN go build -o main cmd/main.go

# Final stage
FROM alpine:3.20
WORKDIR /app

# Menyalin executable dari stage builder
COPY --from=builder /app/main .

# Menjalankan aplikasi
CMD ["./main"]


# FROM golang:alpine3.20
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main .
# CMD ["./main"]