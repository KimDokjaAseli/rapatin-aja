# 📋 Panduan Trello Board - Proyek RapatIn

Dokumen ini berisi daftar tugas, deskripsi, dan checklist yang siap dimasukkan ke dalam Trello Board untuk tim pengembang **RapatIn**. Proyek ini menggunakan arsitektur modular yang memisahkan **Database (MongoDB)**, **Backend (Golang Gin)**, dan **Frontend (Web Statis / Flutter App)**.

---

## 💡 Cara Memasukkan ke Trello secara Cepat
1. **Buat List (Kolom)** di Trello Board Anda:
   * `Database Administrator (DBA)`
   * `Backend Developer`
   * `Frontend Developer`
2. **Impor Card (Kartu) secara Massal**:
   * Salin daftar judul kartu dari bagian **Daftar Ringkasan Kartu** di bawah ini.
   * Klik **"Add a card"** pada list yang sesuai di Trello.
   * Tempel (paste) daftar tersebut, lalu tekan **Enter**.
   * Trello akan bertanya apakah Anda ingin memisahkan setiap baris menjadi kartu baru. Pilih **"Create Cards"**.
3. **Isi Deskripsi & Checklist**:
   * Buka kartu yang telah dibuat, lalu salin **Deskripsi** dan **Checklist** dari bagian detail di bawah ini ke kartu Trello Anda.

---

## 📌 Daftar Ringkasan Kartu (Untuk Salin-Tempel Massal)

### 🗄️ List: Database Administrator (DBA)
```text
DBA-01: Install dan Konfigurasi Server MongoDB Lokal
DBA-02: Pembuatan Database rapatuy_penjaruy dan Koleksi
DBA-03: Penerapan Indeks Database untuk Optimasi Kueri
DBA-04: Seeding Data Uji (User & Meetings) untuk Development
DBA-05: Konfigurasi Kredensial Database & Variabel Lingkungan
```

### ⚙️ List: Backend Developer
```text
BE-01: Inisialisasi Project Go & Konfigurasi Koneksi Database
BE-02: Implementasi API Autentikasi Pengguna (Register & Login)
BE-03: Implementasi API CRUD Notulensi Rapat
BE-04: Implementasi API Manajemen Profil Pengguna
BE-05: Konfigurasi CORS & Penyajian Static Files Frontend
BE-06: Pengujian API Manual & Integrasi Endpoint
```

### 🎨 List: Frontend Developer
```text
FE-01: Implementasi Tema Utama, Styling, dan Responsivitas UI
FE-02: Pembuatan Halaman Autentikasi & Manajemen Sesi Lokal (Session)
FE-03: Pembuatan Dashboard & Tampilan Daftar Rapat dengan Fitur Pencarian
FE-04: Pembuatan Form Modal Tambah, Edit, dan Hapus Rapat
FE-05: Integrasi Halaman Profil Pengguna & Avatar Dinamis (UI-Avatars API)
FE-06: Integrasi API Backend ke UI & Penanganan Status Error (Toast)
```

---

## 🛠️ Detail Kartu Trello (Deskripsi & Checklist)

### 🗄️ DATABASE ADMINISTRATOR (DBA)

#### **DBA-01: Install dan Konfigurasi Server MongoDB Lokal**
* **Label**: `Database` `Setup` `High Priority`
* **Deskripsi**:
  Menyiapkan database server lokal MongoDB agar dapat diakses oleh aplikasi backend selama masa pengembangan (development).
* **Checklist**:
  - [ ] Unduh dan install MongoDB Community Server
  - [ ] Install MongoDB Compass (GUI Client) untuk visualisasi data
  - [ ] Jalankan layanan MongoDB dan pastikan berjalan di port default `27017`
  - [ ] Uji koneksi menggunakan string koneksi: `mongodb://localhost:27017/`

#### **DBA-02: Pembuatan Database rapatuy_penjaruy dan Koleksi**
* **Label**: `Database` `Schema` `High Priority`
* **Deskripsi**:
  Membuat database baru dengan nama `rapatuy_penjaruy` serta mendefinisikan koleksi-koleksi utama yang diperlukan oleh model data aplikasi.
* **Checklist**:
  - [ ] Buat database bernama `rapatuy_penjaruy` di MongoDB
  - [ ] Buat koleksi `users` untuk menyimpan profil pengguna dan kata sandi yang telah di-hash
  - [ ] Buat koleksi `meetings` untuk menyimpan notulensi rapat beserta ID pembuatnya (`userId`)
  - [ ] Verifikasi bahwa koleksi telah berhasil dibuat dan terlihat di MongoDB Compass

#### **DBA-03: Penerapan Indeks Database untuk Optimasi Kueri**
* **Label**: `Database` `Performance` `Medium Priority`
* **Deskripsi**:
  Membuat indeks pada bidang (field) penting untuk menjaga performa kueri pencarian data dan memastikan keunikan data autentikasi.
* **Checklist**:
  - [ ] Terapkan indeks unik (`unique index`) pada field `email` di koleksi `users`
  - [ ] Terapkan indeks pada field `userId` di koleksi `meetings` untuk mempercepat pemuatan rapat per user
  - [ ] Terapkan indeks pada field `title` di koleksi `meetings` untuk pencarian berbasis teks

#### **DBA-04: Seeding Data Uji (User & Meetings) untuk Development**
* **Label**: `Database` `Testing` `Medium Priority`
* **Deskripsi**:
  Memasukkan data awal (dummy data) ke dalam database untuk mempermudah tim backend dan frontend melakukan pengujian fitur login dan tampilan dashboard.
* **Checklist**:
  - [ ] Buat data user uji dengan email `budi@email.com` dan password yang di-hash dengan bcrypt (`rahasia123`)
  - [ ] Buat minimal 2 data rapat contoh yang dihubungkan dengan ID pengguna Budi
  - [ ] Pastikan relasi model referensi ID (`userId` pada dokumen rapat) valid dan sesuai

#### **DBA-05: Konfigurasi Kredensial Database & Variabel Lingkungan**
* **Label**: `Database` `Security` `High Priority`
* **Deskripsi**:
  Menerapkan pengaturan keamanan database untuk staging/produksi dengan membuat user database baru dan merancang file konfigurasi variabel lingkungan (.env).
* **Checklist**:
  - [ ] Buat user database khusus dengan hak akses terbatas (`readWrite`) pada database `rapatuy_penjaruy`
  - [ ] Siapkan konfigurasi variabel lingkungan dengan variabel `MONGO_URL` atau `MONGODB_URL`
  - [ ] Dokumentasikan format koneksi database pada repositori

---

### ⚙️ BACKEND DEVELOPER

#### **BE-01: Inisialisasi Project Go & Konfigurasi Koneksi Database**
* **Label**: `Backend` `Setup` `High Priority`
* **Deskripsi**:
  Menginisialisasi modul Go backend, menginstal library driver MongoDB dan Gin Web Framework, serta membangun file konfigurasi database di `config/db.go`.
* **Checklist**:
  - [ ] Jalankan inisialisasi modul Go (`go mod init rapatln_backend`)
  - [ ] Instal dependensi yang diperlukan (`gin`, `mongo-driver`, `cors`, `crypto/bcrypt`)
  - [ ] Buat file `config/db.go` yang membaca variabel lingkungan `MONGO_URL` / `MONGODB_URL` untuk koneksi database
  - [ ] Buat fungsi ping database di startup untuk memastikan status koneksi sukses

#### **BE-02: Implementasi API Autentikasi Pengguna (Register & Login)**
* **Label**: `Backend` `Feature` `High Priority`
* **Deskripsi**:
  Membangun endpoint untuk pendaftaran akun dan verifikasi masuk dengan enkripsi kata sandi yang aman.
  * **POST /api/register**: Validasi email unik dan enkripsi password menggunakan bcrypt sebelum disimpan.
  * **POST /api/login**: Cocokkan password terenkripsi dengan input pengguna untuk validasi masuk.
* **Checklist**:
  - [ ] Definisikan struct model `User` di `models/user.go`
  - [ ] Buat fungsi `Register` di `controllers/user_controller.go` (gunakan bcrypt untuk hashing password)
  - [ ] Buat fungsi `Login` di `controllers/user_controller.go` (verifikasi hash password)
  - [ ] Hubungkan router group `/api` ke controller Register dan Login di `main.go`

#### **BE-03: Implementasi API CRUD Notulensi Rapat**
* **Label**: `Backend` `Feature` `High Priority`
* **Deskripsi**:
  Membangun endpoint untuk memanipulasi data notulensi rapat dengan filter berdasarkan ID pengguna pembuat rapat.
  * **GET /api/meetings**: Mendapatkan daftar rapat (opsional filter kueri `?userId=`)
  * **GET /api/meetings/:id**: Mendapatkan detail satu rapat
  * **POST /api/meetings**: Membuat rapat baru
  * **PATCH /api/meetings/:id**: Mengubah data rapat secara parsial
  * **DELETE /api/meetings/:id**: Menghapus rapat dari database
* **Checklist**:
  - [ ] Definisikan struct model `Meeting` di `models/meeting.go`
  - [ ] Buat endpoint `GET /meetings` dengan filter `userId` di `controllers/meeting_controller.go`
  - [ ] Buat endpoint `GET /meetings/:id` dan `POST /meetings`
  - [ ] Buat endpoint `PATCH /meetings/:id` untuk update dinamis
  - [ ] Buat endpoint `DELETE /meetings/:id` untuk penghapusan data
  - [ ] Daftarkan semua endpoint di router `main.go`

#### **BE-04: Implementasi API Manajemen Profil Pengguna**
* **Label**: `Backend` `Feature` `Medium Priority`
* **Deskripsi**:
  Membangun endpoint untuk melihat profil pengguna dan mengubah data pribadi (Nama, Jabatan, Bio, Email, Avatar) agar tersinkronisasi dengan database.
  * **GET /api/user/:id**: Mendapatkan profil pengguna
  * **PUT /api/user/:id**: Memperbarui data profil pengguna
* **Checklist**:
  - [ ] Buat fungsi `GetUser` di `controllers/user_controller.go`
  - [ ] Buat fungsi `UpdateUser` di `controllers/user_controller.go` dengan operasi update MongoDB `$set`
  - [ ] Daftarkan rute `/api/user/:id` di router group `/api` dalam `main.go`

#### **BE-05: Konfigurasi CORS & Penyajian Static Files Frontend**
* **Label**: `Backend` `Config` `Medium Priority`
* **Deskripsi**:
  Mengaktifkan CORS (Cross-Origin Resource Sharing) agar client/frontend eksternal dapat melakukan request ke API, serta mengonfigurasi router Gin agar dapat menyajikan file statis frontend langsung dari server Go.
* **Checklist**:
  - [ ] Integrasikan middleware `github.com/gin-contrib/cors` dengan konfigurasi default
  - [ ] Tambahkan konfigurasi penyajian file statis di `main.go` (`r.StaticFile` untuk `/`, `/app.js`, dan `/style.css` yang merujuk ke folder frontend)
  - [ ] Uji apakah port server (default 8080) dapat melayani request dengan lancar

#### **BE-06: Pengujian API Manual & Integrasi Endpoint**
* **Label**: `Backend` `Testing` `High Priority`
* **Deskripsi**:
  Melakukan verifikasi seluruh API endpoint yang telah dibuat menggunakan POSTMAN, Insomnia, atau Thunder Client untuk memastikan respons JSON dan status HTTP sesuai spesifikasi.
* **Checklist**:
  - [ ] Uji skenario registrasi dengan input valid dan input email duplikat
  - [ ] Uji skenario login dengan password benar dan salah (pastikan JSON response membawa data pengguna secara lengkap tanpa menampilkan password)
  - [ ] Uji CRUD Rapat: input data baru, update data, dan hapus data (pastikan perubahan langsung tercermin di database MongoDB)
  - [ ] Pastikan respons error (status 400/401/404/500) mengembalikan pesan JSON yang deskriptif

---

### 🎨 FRONTEND DEVELOPER

#### **FE-01: Implementasi Tema Utama, Styling, dan Responsivitas UI**
* **Label**: `Frontend` `UI-UX` `High Priority`
* **Deskripsi**:
  Membangun pondasi tampilan aplikasi dengan skema warna Maroon modern (atau tema premium yang disepakati), menetapkan tipografi modern, serta membuat layout yang responsif untuk tampilan desktop dan mobile.
* **Checklist**:
  * **Jika Web Statis**:
    - [ ] Buat struktur HTML utama di `fe/index.html` (Header, Sidebar, Area Content, Modals)
    - [ ] Buat desain CSS di `fe/style.css` dengan variabel warna, efek glassmorphism, dan hover transition
    - [ ] Pastikan layout responsif menggunakan Grid/Flexbox
  * **Jika Flutter (Mobile)**:
    - [ ] Atur tema warna aplikasi di `app_theme.dart` (menggunakan warna maroon)
    - [ ] Desain custom widget tombol, input textfield dengan border bulat, dan kartu informasi

#### **FE-02: Pembuatan Halaman Autentikasi & Manajemen Sesi Lokal (Session)**
* **Label**: `Frontend` `Feature` `High Priority`
* **Deskripsi**:
  Membangun tampilan form Login dan Register, menangani submit data ke API backend, serta menyimpan data sesi ke penyimpanan lokal browser / perangkat.
* **Checklist**:
  - [ ] Buat tampilan UI Form Login dan Register (dengan validasi kolom kosong)
  - [ ] Implementasikan fungsi pengiriman data ke backend `/api/login` dan `/api/register`
  - [ ] Simpan data user ke `localStorage` (untuk Web, dengan key `rapatin_user`) atau Secure Storage (untuk Flutter) setelah login berhasil
  - [ ] Buat fungsi auto-login (cek session di lokal penyimpanan saat aplikasi pertama kali dibuka)
  - [ ] Buat fungsi Logout untuk menghapus data sesi dan mengarahkan kembali pengguna ke halaman login

#### **FE-03: Pembuatan Dashboard & Tampilan Daftar Rapat dengan Fitur Pencarian**
* **Label**: `Frontend` `Feature` `High Priority`
* **Deskripsi**:
  Membangun halaman utama (Dashboard/Home) yang memuat daftar rapat yang telah dibuat oleh pengguna bersangkutan dan menyediakan input pencarian dinamis (real-time search).
* **Checklist**:
  - [ ] Ambil data rapat dari `/api/meetings?userId={id}` saat dashboard dimuat
  - [ ] Tampilkan ringkasan rapat dalam format kartu rapat (`MeetingCard`) secara berurutan
  - [ ] Tampilkan maksimal 4 rapat terbaru di halaman Beranda/Home dan seluruh rapat di tab/halaman Rapat
  - [ ] Implementasikan kolom pencarian di bagian atas untuk memfilter daftar rapat berdasarkan judul atau lokasi rapat secara real-time
  - [ ] Tampilkan status loading (`CircularProgressIndicator` / Skeleton loader) saat data sedang dimuat

#### **FE-04: Pembuatan Form Modal Tambah, Edit, dan Hapus Rapat**
* **Label**: `Frontend` `Feature` `High Priority`
* **Deskripsi**:
  Membuat formulir interaktif (dalam bentuk Modal Popup atau Halaman Form baru) untuk menambah, mengubah, dan menghapus catatan rapat.
* **Checklist**:
  - [ ] Buat form input untuk Judul, Tanggal (menggunakan date picker), Waktu, Lokasi, dan Deskripsi Rapat
  - [ ] Hubungkan tombol **Buat Rapat** / **Simpan** untuk mengirimkan data ke endpoint `POST /api/meetings`
  - [ ] Buat form edit yang terisi otomatis dengan data rapat sebelumnya dan kirim pembaruan ke `PATCH /api/meetings/:id`
  - [ ] Tambahkan dialog konfirmasi konseptual sebelum pengguna menghapus rapat (`DELETE /api/meetings/:id`)
  - [ ] Pastikan daftar rapat di UI ter-update secara instan tanpa perlu refresh browser / memuat ulang paksa halaman

#### **FE-05: Integrasi Halaman Profil Pengguna & Avatar Dinamis (UI-Avatars API)**
* **Label**: `Frontend` `Feature` `Medium Priority`
* **Deskripsi**:
  Membangun tampilan halaman profil yang menampilkan data detail pengguna, total statistik rapat yang telah dibuat, form untuk mengubah profil, dan mengintegrasikan pembuatan avatar inisial otomatis.
* **Checklist**:
  - [ ] Tampilkan detail Nama, Email, Jabatan, dan Bio Pengguna
  - [ ] Buat form edit profil untuk memperbarui data ke API backend `PUT /api/user/:id`
  - [ ] Integrasikan URL dynamic avatar dari `ui-avatars.com` berdasarkan nama lengkap pengguna (misal: `https://ui-avatars.com/api/?name=Budi+Santoso`)
  - [ ] Hitung secara dinamis total rapat yang telah dibuat oleh pengguna untuk ditampilkan di bagian statistik profil

#### **FE-06: Integrasi API Backend ke UI & Penanganan Status Error (Toast)**
* **Label**: `Frontend` `Integration` `High Priority`
* **Deskripsi**:
  Menghubungkan frontend secara penuh dengan backend, mengatur penanganan error API (misalnya email sudah terdaftar, login gagal, server down), dan menampilkan notifikasi toast (feedback) yang menarik ke pengguna.
* **Checklist**:
  - [ ] Pastikan endpoint server API (URL dan Port) diarahkan dengan benar via konfigurasi global (base URL)
  - [ ] Tampilkan notifikasi toast sukses/gagal yang menarik untuk setiap aksi (Register, Login, CRUD Rapat, Edit Profil)
  - [ ] Handle error request (tampilkan pesan kesalahan yang dikirimkan oleh backend JSON, misal: `"Invalid email or password"`)
  - [ ] Lakukan uji coba skenario end-to-end (E2E) dari pendaftaran akun hingga penghapusan rapat untuk memvalidasi kelancaran integrasi
