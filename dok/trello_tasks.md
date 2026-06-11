# 📋 Panduan Trello Board - Proyek RapatIn

Dokumen ini berisi daftar tugas, deskripsi, dan checklist yang siap dimasukkan ke Trello Board tim **RapatIn**. Proyek ini dibagi menjadi tiga bagian: **Database (MongoDB)**, **Backend (Golang Gin)**, dan **Frontend (Web Statis / Flutter App)**.

---

## 💡 Cara Cepat Memasukkan Tugas ke Trello
1. **Buat Kolom (List)** di Trello Board Anda:
   * `Database Administrator (DBA)`
   * `Backend Developer`
   * `Frontend Developer`
2. **Impor Kartu (Card) secara Massal**:
   * Salin daftar judul kartu dari bagian **Daftar Ringkasan Kartu** di bawah.
   * Klik **"Add a card"** pada kolom yang sesuai di Trello.
   * Tempel (paste) daftar judul tersebut, lalu tekan **Enter**.
   * Pilih **"Create Cards"** ketika Trello bertanya untuk memisahkan setiap baris menjadi kartu baru.
3. **Isi Deskripsi & Checklist**:
   * Buka kartu yang sudah dibuat, lalu salin **Deskripsi** dan **Checklist** dari detail di bawah ke kartu Trello Anda.

---

## 📌 Daftar Ringkasan Kartu (Untuk Salin-Tempel Massal)

### 🗄️ List: Database Administrator (DBA)
```text
DBA-01: Install dan Jalankan MongoDB Lokal
DBA-02: Buat Database rapatuy_penjaruy dan Tabel/Koleksi
DBA-03: Pasang Indeks Database agar Pencarian Lebih Cepat
DBA-04: Isi Data Contoh (User & Rapat) untuk Uji Coba
DBA-05: Atur Keamanan Database & File .env
```

### ⚙️ List: Backend Developer
```text
BE-01: Setup Project Go & Hubungkan ke Database
BE-02: Buat API Autentikasi (Register & Login)
BE-03: Buat API Mengelola Rapat (CRUD)
BE-04: Buat API Profil Pengguna
BE-05: Atur CORS & Sajikan File Frontend dari Backend
BE-06: Tes Semua API secara Manual
```

### 🎨 List: Frontend Developer
```text
FE-01: Desain Tampilan Utama (Tema Maroon & Responsif)
FE-02: Buat Halaman Login/Register & Simpan Sesi Login
FE-03: Buat Dashboard & Daftar Rapat dengan Kolom Pencarian
FE-04: Buat Form Modal (Tambah, Edit, Hapus Rapat)
FE-05: Buat Halaman Profil & Foto Profil Otomatis (UI-Avatars)
FE-06: Hubungkan UI Frontend dengan API Backend & Atur Notifikasi (Toast)
```

---

## 🛠️ Detail Kartu Trello (Deskripsi & Checklist)

### 🗄️ DATABASE ADMINISTRATOR (DBA)

#### **DBA-01: Install dan Jalankan MongoDB Lokal**
* **Label**: `Database` `Setup` `High Priority`
* **Deskripsi**:
  Menginstal dan menjalankan database MongoDB di komputer/laptop lokal agar bisa dipakai oleh backend selama masa pembuatan aplikasi.
* **Checklist**:
  - [ ] Download dan instal MongoDB Community Server
  - [ ] Instal MongoDB Compass (aplikasi visual untuk melihat isi database)
  - [ ] Pastikan MongoDB berjalan di port bawaan `27017`
  - [ ] Tes koneksi ke database dengan alamat `mongodb://localhost:27017/`

#### **DBA-02: Buat Database rapatuy_penjaruy dan Tabel/Koleksi**
* **Label**: `Database` `Schema` `High Priority`
* **Deskripsi**:
  Membuat database baru bernama `rapatuy_penjaruy` beserta tabel-tabel (koleksi) utama yang dibutuhkan aplikasi.
* **Checklist**:
  - [ ] Buat database baru bernama `rapatuy_penjaruy`
  - [ ] Buat koleksi `users` (untuk menyimpan data akun dan password)
  - [ ] Buat koleksi `meetings` (untuk menyimpan data rapat dan ID pembuatnya)
  - [ ] Pastikan kedua koleksi sudah muncul di MongoDB Compass

#### **DBA-03: Pasang Indeks Database agar Pencarian Lebih Cepat**
* **Label**: `Database` `Performance` `Medium Priority`
* **Deskripsi**:
  Membuat indeks pada kolom penting agar pencarian data lebih cepat dan mencegah email kembar di database.
* **Checklist**:
  - [ ] Buat indeks unik untuk kolom `email` di tabel `users` (email tidak boleh sama)
  - [ ] Buat indeks pada kolom `userId` di tabel `meetings` agar loading rapat per user lebih cepat
  - [ ] Buat indeks pencarian teks pada kolom `title` di tabel `meetings` untuk fitur pencarian

#### **DBA-04: Isi Data Contoh (User & Rapat) untuk Uji Coba**
* **Label**: `Database` `Testing` `Medium Priority`
* **Deskripsi**:
  Mengisi database dengan data contoh (dummy data) agar backend dan frontend bisa langsung menguji fitur login dan dashboard.
* **Checklist**:
  - [ ] Buat data user contoh: email `budi@email.com` dan password `rahasia123` (di-hash menggunakan bcrypt)
  - [ ] Buat minimal 2 data rapat contoh yang terhubung ke user Budi
  - [ ] Pastikan ID user (`userId`) pada data rapat sudah sesuai dengan ID user Budi

#### **DBA-05: Atur Keamanan Database & File .env**
* **Label**: `Database` `Security` `High Priority`
* **Deskripsi**:
  Mengatur keamanan akses database dan menyiapkan file konfigurasi variabel lingkungan (`.env`).
* **Checklist**:
  - [ ] Buat akun admin database khusus dengan hak akses `readWrite` pada database `rapatuy_penjaruy`
  - [ ] Siapkan template file `.env` yang berisi variabel `MONGO_URL` atau `MONGODB_URL`
  - [ ] Tulis dokumentasi singkat tentang cara koneksi database di file README proyek

---

### ⚙️ BACKEND DEVELOPER

#### **BE-01: Setup Project Go & Hubungkan ke Database**
* **Label**: `Backend` `Setup` `High Priority`
* **Deskripsi**:
  Membuat proyek Go baru, menginstal library/package yang dibutuhkan, dan membuat file konfigurasi database di `config/db.go`.
* **Checklist**:
  - [ ] Jalankan perintah `go mod init` (misal: `go mod init rapatin_backend`)
  - [ ] Instal package yang dibutuhkan: `gin`, `mongo-driver`, `cors`, dan `bcrypt`
  - [ ] Buat file `config/db.go` untuk membaca alamat MongoDB dari file `.env`
  - [ ] Buat fungsi tes koneksi (ping database) saat aplikasi backend pertama kali dinyalakan

#### **BE-02: Buat API Autentikasi (Register & Login)**
* **Label**: `Backend` `Feature` `High Priority`
* **Deskripsi**:
  Membuat API untuk pendaftaran akun baru (Register) dan masuk log (Login).
  * **POST /api/register**: Menyimpan user baru dengan password yang sudah di-hash (dienkripsi) menggunakan bcrypt.
  * **POST /api/login**: Memvalidasi password user untuk masuk aplikasi.
* **Checklist**:
  - [ ] Buat model data `User` di folder `models/user.go`
  - [ ] Buat fungsi `Register` di controller (password di-hash dengan bcrypt sebelum disimpan)
  - [ ] Buat fungsi `Login` di controller (mencocokkan password input dengan database)
  - [ ] Daftarkan route `/api/register` dan `/api/login` di file `main.go`

#### **BE-03: Buat API Mengelola Rapat (CRUD)**
* **Label**: `Backend` `Feature` `High Priority`
* **Deskripsi**:
  Membuat API untuk membuat, melihat, mengedit, dan menghapus data rapat berdasarkan user yang membuatnya.
  * **GET /api/meetings**: Mengambil daftar rapat (bisa difilter menggunakan query `?userId=`)
  * **GET /api/meetings/:id**: Melihat detail satu rapat
  * **POST /api/meetings**: Membuat rapat baru
  * **PATCH /api/meetings/:id**: Mengubah sebagian data rapat
  * **DELETE /api/meetings/:id**: Menghapus data rapat
* **Checklist**:
  - [ ] Buat model data `Meeting` di folder `models/meeting.go`
  - [ ] Buat API untuk mengambil daftar rapat (bisa difilter dengan `userId`)
  - [ ] Buat API untuk melihat detail rapat dan menambah rapat baru
  - [ ] Buat API untuk edit rapat (PATCH) dan hapus rapat (DELETE)
  - [ ] Daftarkan semua route rapat tersebut di file `main.go`

#### **BE-04: Buat API Profil Pengguna**
* **Label**: `Backend` `Feature` `Medium Priority`
* **Deskripsi**:
  Membuat API untuk menampilkan profil user dan mengubah data diri (Nama, Jabatan, Bio, Email, Avatar).
  * **GET /api/user/:id**: Mengambil data profil user
  * **PUT /api/user/:id**: Mengupdate data profil user di database
* **Checklist**:
  - [ ] Buat fungsi ambil user di controller
  - [ ] Buat fungsi edit user di controller (menggunakan update `$set` MongoDB)
  - [ ] Daftarkan route `/api/user/:id` di file `main.go`

#### **BE-05: Atur CORS & Sajikan File Frontend dari Backend**
* **Label**: `Backend` `Config` `Medium Priority`
* **Deskripsi**:
  Mengaktifkan CORS agar frontend eksternal bisa mengakses API, serta mengatur agar backend bisa langsung menampilkan file HTML/CSS/JS frontend (jika menggunakan web statis).
* **Checklist**:
  - [ ] Pasang middleware CORS (menggunakan `github.com/gin-contrib/cors`)
  - [ ] Atur route statis di `main.go` agar server Go bisa melayani file frontend langsung
  - [ ] Uji coba apakah server backend (port 8080) berjalan lancar saat diakses

#### **BE-06: Tes Semua API secara Manual**
* **Label**: `Backend` `Testing` `High Priority`
* **Deskripsi**:
  Menguji seluruh API yang telah dibuat menggunakan Postman, Insomnia, atau Thunder Client untuk memastikan respon data dan status kode HTTP sudah benar.
* **Checklist**:
  - [ ] Uji register user baru dan tes jika email kembar (harus error)
  - [ ] Uji login dengan password yang benar dan salah (pastikan data password asli tidak ikut dikirim di respon JSON)
  - [ ] Uji tambah, edit, dan hapus rapat, lalu pastikan datanya langsung berubah di MongoDB
  - [ ] Pastikan jika terjadi error, API mengembalikan pesan error yang jelas dalam format JSON

---

### 🎨 FRONTEND DEVELOPER

#### **FE-01: Desain Tampilan Utama (Tema Maroon & Responsif)**
* **Label**: `Frontend` `UI-UX` `High Priority`
* **Deskripsi**:
  Membuat desain dasar aplikasi dengan warna Maroon modern, font yang bagus, dan tampilan yang rapi saat dibuka di HP maupun laptop.
* **Checklist**:
  * **Jika Web Statis (HTML/CSS)**:
    - [ ] Buat kerangka HTML utama (`fe/index.html`): Header, Sidebar, Area Konten, dan Modal Popup
    - [ ] Desain CSS (`fe/style.css`): atur warna maroon, efek blur (glassmorphism), dan transisi tombol yang halus
    - [ ] Pastikan tampilan rapi di layar HP (responsif menggunakan Flexbox/Grid)
  * **Jika Flutter (Mobile)**:
    - [ ] Atur tema warna maroon di file `app_theme.dart`
    - [ ] Buat komponen (widget) custom untuk tombol, input text, dan kartu informasi

#### **FE-02: Buat Halaman Login/Register & Simpan Sesi Login**
* **Label**: `Frontend` `Feature` `High Priority`
* **Deskripsi**:
  Membuat formulir Login & Register, menghubungkannya ke API backend, serta menyimpan data login agar user tidak perlu login ulang saat membuka aplikasi.
* **Checklist**:
  - [ ] Buat tampilan halaman Login dan Register
  - [ ] Hubungkan form ke API `/api/login` dan `/api/register`
  - [ ] Simpan data user ke `localStorage` (untuk Web) atau Secure Storage (untuk Flutter) setelah login berhasil
  - [ ] Buat fitur cek sesi login otomatis saat aplikasi pertama kali dibuka (auto-login)
  - [ ] Buat tombol Logout untuk menghapus data sesi dan kembali ke halaman Login

#### **FE-03: Buat Dashboard & Daftar Rapat dengan Kolom Pencarian**
* **Label**: `Frontend` `Feature` `High Priority`
* **Deskripsi**:
  Membuat halaman utama (Dashboard) untuk menampilkan daftar rapat milik user dan menambahkan kolom pencarian langsung (real-time).
* **Checklist**:
  - [ ] Panggil API `/api/meetings?userId={id}` untuk mengambil data rapat saat dashboard dibuka
  - [ ] Tampilkan daftar rapat menggunakan kartu desain (card) secara berurutan
  - [ ] Tampilkan maksimal 4 rapat terbaru di halaman Beranda, dan tampilkan semua rapat di halaman/tab Rapat
  - [ ] Buat kolom pencarian di bagian atas untuk menyaring rapat berdasarkan judul atau lokasi
  - [ ] Tampilkan animasi loading (loading spinner atau skeleton card) saat data sedang dimuat

#### **FE-04: Buat Form Modal (Tambah, Edit, Hapus Rapat)**
* **Label**: `Frontend` `Feature` `High Priority`
* **Deskripsi**:
  Membuat modal popup atau formulir untuk menambah, mengedit, dan menghapus data rapat.
* **Checklist**:
  - [ ] Buat form input: Judul, Tanggal (menggunakan kalender/date picker), Waktu, Lokasi, dan Deskripsi
  - [ ] Hubungkan tombol simpan ke API tambah rapat (`POST /api/meetings`)
  - [ ] Buat form edit yang otomatis terisi data rapat lama dan hubungkan ke API edit (`PATCH /api/meetings/:id`)
  - [ ] Buat kotak konfirmasi (dialog) sebelum user menghapus rapat (`DELETE /api/meetings/:id`)
  - [ ] Pastikan daftar rapat langsung ter-update di layar setelah ditambah/diedit/dihapus (tanpa perlu reload browser manual)

#### **FE-05: Buat Halaman Profil & Foto Profil Otomatis (UI-Avatars)**
* **Label**: `Frontend` `Feature` `Medium Priority`
* **Deskripsi**:
  Membuat halaman profil yang menampilkan data diri user, jumlah rapat yang sudah dibuat, form edit profil, dan foto profil inisial otomatis.
* **Checklist**:
  - [ ] Tampilkan detail Nama, Email, Jabatan, dan Bio user
  - [ ] Buat form edit profil untuk menyimpan perubahan ke API `/api/user/:id`
  - [ ] Hubungkan foto profil dengan API `ui-avatars.com` berdasarkan nama user (misal: `https://ui-avatars.com/api/?name=Budi+Santoso`)
  - [ ] Hitung total rapat milik user dan tampilkan sebagai angka statistik di profil

#### **FE-06: Hubungkan UI Frontend dengan API Backend & Atur Notifikasi (Toast)**
* **Label**: `Frontend` `Integration` `High Priority`
* **Deskripsi**:
  Menghubungkan frontend ke backend secara penuh, menangani pesan kesalahan dari server, dan memunculkan notifikasi pop-up (toast) yang menarik.
* **Checklist**:
  - [ ] Atur alamat base URL API backend agar terpusat di satu file konfigurasi
  - [ ] Tampilkan notifikasi pop-up (toast) sukses/gagal untuk setiap aksi (login, register, tambah/edit/hapus rapat, edit profil)
  - [ ] Tangani error dari API (misal tampilkan pesan "Email sudah terdaftar" atau "Password salah") agar mudah dibaca user
  - [ ] Lakukan uji coba alur aplikasi dari awal (daftar akun) sampai akhir (hapus rapat) untuk memastikan semua berjalan lancar
