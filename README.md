# Simple Bank Rest API

Ini adalah rest API sederhana yang dibuat menggunakan bahasa pemrograman Golang. Rest API ini digunakan untuk
menyediakan layanan untuk customer dapat melakukan pembayaran kepada merchant.

## Endpoint

### 1. **Login**

**URL:** `/auth/login`  
**Method:** `POST`  
**Deskripsi:** Autentikasi pengguna dan memberikan token sesi.

**Request Body:**

```json
{
  "username": "username",
  "password": "password"
}
```
### 2. **Logout**

**URL:** `/auth/logout`  
**Method:** `POST`
**Authorization:** `Basic <Token>`

**Deskripsi:** Logout pengguna dengan menghapus token sesi

### 3. **Payment**

**URL:** `/payments`  
**Method:** `POST`
**Authorization:** `Basic <Token>`

**Deskripsi:** Customer melakukan pembayaran ke merchant

**Request Body:**

```json
{
  "amount": 10000
}
```

### 4. **Customer**

**URL:** `/customers`  
**Method:** `GET`
**Authorization:** `Basic <Token>`

**Deskripsi:** Menampilkan daftar customer yang terdaftar

## Fitur

- **Autentikasi Pengguna**
    - Login
    - Logout
- **Manajemen Pelanggan**
    - Melihat daftar pelanggan
- **Pembayaran**
    - Transfer pembayaran dari pelanggan ke merchant

---

## Instalasi
### 1. **Melakukan Clone Repository**
```bash
  git clone https://github.com/Refanz/simple-bank-rest-api
  cd simple-bank-rest-api/app
```

### 2. **Menjalankan Server**
```bash
  go run main.go
```


### 3. **Mengakses API**
Server akan berjalan di `http://localhost:8080`

## Catatan
- Endpoint autentikasi menggunakan token basic yang didapatkan dari hasil encoding username dan password. Pastikan token dimasukkan dalam header `Authorization` untuk endpoint yang membutuhkan autentikasi
- API ini belum terhubung ke database. Data masih disimpan dalam file `data.json`.