# RestApi-Golang

Rest API sederhana menggunakan Golang dengan **Chi Router** dan database **MongoDB**.

## 🚀 Fitur
- CRUD (Create, Read, Update, Delete) untuk data tertentu.
- Menggunakan **Chi Router** sebagai HTTP router.
- Menggunakan **MongoDB** sebagai database.
- Menggunakan **dotenv** untuk manajemen konfigurasi.
- Menggunakan **SCRAM-SHA-1** untuk otentikasi MongoDB.

## 📦 Instalasi

### 1. Clone Repository
```sh
git clone https://github.com/Dito-7/RestApi-Golang.git
cd RestApi-Golang
```

### 2. Install Dependencies
Pastikan Anda sudah menginstal Golang, lalu jalankan:
```sh
go mod tidy
```

### 3. Konfigurasi Environment
Buat file `.env` dan sesuaikan dengan konfigurasi MongoDB Anda:
```env
MONGO_URI=""
MONGO_DBNAME=""
MONGO_COLLECTION_NAME=""
```

### 4. Jalankan Server
```sh
go run main.go
```
Server akan berjalan di `http://localhost:4444`

## 🛠 Teknologi yang Digunakan
- [Golang](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- [MongoDB](https://www.mongodb.com/)
- [godotenv](https://github.com/joho/godotenv)

## 📜 Lisensi
MIT License - bebas digunakan dan dikembangkan!

---
Jika ada pertanyaan atau ingin kontribusi, jangan ragu untuk membuat **issue** atau **pull request**! 😊
