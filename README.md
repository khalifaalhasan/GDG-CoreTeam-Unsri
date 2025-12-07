# 🚀 GDG UNSRI Core API

Backend server untuk manajemen **Event** dan **User**, dibangun
menggunakan **Go (Golang)** dengan fokus pada keamanan **Role-Based
Access Control (RBAC)** menggunakan **Firebase Authentication**.

------------------------------------------------------------------------

## 📦 Stack Teknologi

  Komponen             Teknologi
  -------------------- -----------------------------
  Bahasa Pemrograman   Go (Golang) 1.21+
  Web Framework        Gin Gonic
  Database             Firebase Firestore
  Autentikasi          Firebase Authentication
  Arsitektur           Modular (Router per Module)

------------------------------------------------------------------------

## 🛠️ Prerequisites

Sebelum menjalankan backend ini, pastikan Anda memiliki:

-   Go Lang versi **1.21 atau lebih tinggi**
-   Postman atau tool pengujian API lainnya
-   Akun Firebase Project dengan:
    -   Firebase Auth aktif\
    -   Firestore aktif\
-   File **serviceAccountKey.json** valid, diletakkan dalam folder
    `config/`

------------------------------------------------------------------------

## ▶️ Cara Menjalankan Server

1.  Pastikan file `config/serviceAccountKey.json` tersedia.
2.  Buka terminal di root project.
3.  Jalankan server:

``` bash
go run cmd/api/main.go
```

Server akan berjalan pada:

    http://localhost:8081

------------------------------------------------------------------------

## 🔑 Panduan Pengujian Security (RBAC)

Semua endpoint **(kecuali public)** dilindungi oleh Firebase ID Token
(JWT).

### ✅ Langkah 1 --- Mendapatkan ID Token (Login)

Gunakan Firebase Identity Toolkit:

**Endpoint:**

    POST https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=[YOUR_FIREBASE_WEB_API_KEY]

**Ambil nilai:** - `idToken` → digunakan sebagai **Bearer Token**

------------------------------------------------------------------------

### 🔐 Langkah 2 --- Uji Hak Akses

Gunakan header berikut pada setiap request yang diproteksi:

    Authorization: Bearer <TOKEN_ANDA>

### 🧪 Skenario Uji RBAC

  Token Digunakan   Endpoint      Method   Ekspektasi      Keterangan
  ----------------- ------------- -------- --------------- ------------------------
  None              `/events`     GET      200 OK          Public
  Member/Admin      `/users/me`   GET      200 OK          AuthMiddleware
  Member            `/events`     POST     403 Forbidden   Ditolak RoleMiddleware
  Admin             `/users`      GET      200 OK          Diizinkan

------------------------------------------------------------------------

## 📋 Referensi Endpoint API

Base URL:

    http://localhost:8081/api/v1

------------------------------------------------------------------------

## 1️⃣ User & Authentication Routes

  -------------------------------------------------------------------------
  Endpoint                   Method    Akses                 Deskripsi
  -------------------------- --------- --------------------- --------------
  `/users/register`          POST      Wajib Login           Sinkronisasi
                                                             profil user ke
                                                             Firestore

  `/users/me`                GET       Wajib Login           Ambil data
                                                             user yang
                                                             sedang login

  `/users/me`                PUT       Wajib Login           Update nama
                                                             user

  `/users`                   GET       Admin Only            List semua
                                                             member

  `/users/:id`               GET       Admin Only            Detail user

  `/users/:id/job`           POST      Admin Only            Berikan job
                                                             desk
  
------------------------------------------------------------------------

## 2️⃣ Event Routes
-------------------------------------------------------------------------
  Endpoint                   Method    Akses                 Deskripsi
  -------------------------- --------- --------------------- --------------
  `/events`                GET         Publik           Mendapatkan semua event

  `/events/:id`            GET         Publik           Mendapatkan detail event

  `/events`                POST        Admin Only       Membuat event baru

  `/events/:id`            PUT         Admin Only       Mengubah data event

  `/events/:id`            DELETE      Admin Only       Menghapus event


------------------------------------------------------------------------

## 📝 Catatan untuk Tester

-   Pastikan menggunakan **Bearer Token yang valid** untuk setiap
    endpoint protected.
-   Jika role bukan Admin, request ke endpoint Admin akan **otomatis
    ditolak**.
-   Gunakan Postman Collection untuk mempercepat pengujian (jika
    tersedia).

------------------------------------------------------------------------

## 📚 Selesai

README ini dirancang agar tester dapat memahami alur login, token, dan
RBAC dengan cepat serta langsung melakukan pengujian API tanpa
kebingungan.
