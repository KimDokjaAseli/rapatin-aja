# 📘 Dokumentasi Fitur & Panduan Pengujian RapatIn

Dokumen ini menjelaskan daftar fitur yang tersedia di aplikasi **RapatIn** beserta skenario uji (test cases) untuk memvalidasi kelancaran integrasi antara frontend dan backend.

---

## 1. Fitur Utama Aplikasi

### A. Autentikasi Pengguna
* **Daftar Akun (Register)**: Mendaftarkan nama lengkap, email, dan kata sandi. Kata sandi akan di-hash secara aman menggunakan bcrypt di backend.
* **Masuk Akun (Login)**: Verifikasi email dan kata sandi dengan database MongoDB.
* **Manajemen Sesi (Session Management)**: Sesi disimpan secara lokal pada browser (`localStorage` dengan kunci `rapatin_user`) agar pengguna tidak perlu masuk kembali saat menyegarkan halaman.
* **Keluar Akun (Logout)**: Menghapus data sesi lokal dan mengarahkan kembali ke halaman autentikasi.

### B. Manajemen Notulensi Rapat
* **Tambah Rapat**: Membuat entri rapat baru dengan atribut Judul, Tanggal, Waktu, Lokasi, dan Deskripsi/Notulensi. Rapat otomatis dihubungkan dengan ID pengguna yang sedang masuk.
* **Daftar Rapat**: Menampilkan daftar rapat yang diurutkan di halaman Beranda (maksimal 4 rapat terbaru) dan halaman Rapat (seluruh rapat).
* **Edit Rapat**: Memperbarui informasi rapat yang sudah disimpan melalui dialog popup.
* **Hapus Rapat**: Menghapus entri notulensi rapat secara permanen dari database MongoDB setelah konfirmasi pengguna.
* **Pencarian Rapat**: Menyaring daftar rapat secara dinamis (real-time) berdasarkan judul atau lokasi rapat melalui bar pencarian di bagian atas.

### C. Profil Pengguna
* **Tampilan Profil**: Menampilkan detail profil berupa Nama Lengkap, Jabatan/Posisi, Bio Singkat, Alamat Email, serta statistik total rapat yang telah dibuat.
* **Ubah Profil**: Memungkinkan pengguna untuk mengedit seluruh informasi profil (Nama, Jabatan, Bio, dan Email) secara langsung yang langsung disimpan ke MongoDB dan memperbarui sesi lokal.
* **Avatar Dinamis**: Menghasilkan gambar profil inisial otomatis menggunakan API `ui-avatars.com` berdasarkan nama lengkap pengguna.

---

## 2. Panduan Pengujian (Skenario Uji)

Jalankan backend dengan perintah `go run main.go` atau jalankan `server.exe`, lalu buka aplikasi di browser (misalnya di `http://localhost:8080/`).

### Skenario Uji 1: Registrasi Pengguna Baru
1. Buka aplikasi. Halaman masuk (Login) akan muncul.
2. Klik tautan **Daftar** di bagian bawah.
3. Isi kolom:
   * **Nama Lengkap**: `Budi Santoso`
   * **Email**: `budi@email.com`
   * **Kata Sandi**: `rahasia123`
4. Klik tombol **Daftar**.
5. **Ekspektasi Hasil**: Muncul pemberitahuan toast sukses `"Pendaftaran berhasil! Silakan masuk."` dan form otomatis beralih kembali ke halaman masuk.

### Skenario Uji 2: Autentikasi Masuk & Keluar Sesi
1. Pada halaman masuk, masukkan email `budi@email.com` dan kata sandi `rahasia123`.
2. Klik tombol **Masuk**.
3. **Ekspektasi Hasil**: 
   * Muncul toast `"Selamat datang kembali!"`.
   * Form login tertutup/tersembunyi.
   * Nama panggilan (`Budi`) dan avatar inisial muncul di bagian pojok kanan atas header.
   * Cek Developer Tools Browser (`F12` -> Application -> Local Storage): Pastikan terdapat item `rapatin_user` berisi JSON data profil pengguna.
4. Klik tombol **Keluar** (di sidebar sebelah kiri atau menu bawah).
5. **Ekspektasi Hasil**: Halaman login muncul kembali dan item `rapatin_user` pada `localStorage` terhapus.

### Skenario Uji 3: Pembuatan Rapat Baru
1. Pastikan Anda telah masuk log.
2. Klik tombol **Buat Rapat** (tombol melayang merah di kanan bawah atau tombol merah di bilah navigasi bawah ponsel).
3. Isi formulir rapat:
   * **Judul Rapat**: `Rapat Evaluasi Sprint 1`
   * **Tanggal**: Pilih tanggal saat ini
   * **Waktu**: Isi waktu luang
   * **Lokasi**: `Ruang Meeting Utama`
   * **Deskripsi / Notulensi**: `Membahas performa aplikasi, integrasi database, dan bug fix.`
4. Klik **Simpan**.
5. **Ekspektasi Hasil**: 
   * Form tertutup.
   * Rapat baru ditambahkan di Beranda dan menu Rapat.
   * Muncul notifikasi toast `"Notulensi rapat baru berhasil dibuat"`.

### Skenario Uji 4: Pencarian Rapat
1. Buat minimal dua rapat dengan judul/lokasi berbeda (misalnya: `Rapat Evaluasi Sprint` di `Ruang Utama` dan `Brainstorming UI` di `Cafetaria`).
2. Pada kolom pencarian di bagian header, ketik kata kunci `UI` atau `Cafetaria`.
3. **Ekspektasi Hasil**: Daftar kartu rapat otomatis menyaring dan hanya menampilkan `Brainstorming UI`. Kosongkan pencarian untuk menampilkan kembali seluruh rapat.

### Skenario Uji 5: Edit Rapat
1. Klik salah satu kartu rapat yang telah dibuat.
2. Form edit akan terbuka dengan data yang terisi otomatis.
3. Ubah judul rapat menjadi `Rapat Evaluasi Sprint 1 - REVISI` dan lokasi menjadi `Ruang Direksi`.
4. Klik **Simpan**.
5. **Ekspektasi Hasil**:
   * Muncul toast `"Notulensi rapat berhasil diperbarui"`.
   * Judul dan lokasi rapat pada kartu rapat langsung berubah secara dinamis di layar tanpa perlu menyegarkan (refresh) browser.

### Skenario Uji 6: Edit Profil Pengguna
1. Buka menu **Profil** dari sidebar kiri atau navigasi bawah.
2. Ubah kolom formulir profil:
   * **Jabatan / Posisi**: `Lead Developer`
   * **Bio Singkat**: `Fokus pada stabilitas sistem.`
3. Klik tombol **Simpan Perubahan**.
4. **Ekspektasi Hasil**:
   * Muncul toast `"Profil berhasil diperbarui"`.
   * Teks jabatan di atas profil berubah menjadi `Lead Developer`.
   * Teks bio berubah menjadi `"Fokus pada stabilitas sistem."`.
   * Inisial nama di avatar dan nama di pojok kanan atas header tetap sinkron dengan nama saat ini.
   * Lakukan penyegaran (refresh) browser, data profil baru Anda tetap dipertahankan.

### Skenario Uji 7: Hapus Rapat
1. Pada kartu rapat di halaman Beranda atau Rapat, klik tombol ikon sampah berwarna merah.
2. Dialog konfirmasi browser akan muncul: `"Apakah Anda yakin ingin menghapus notulensi ini?"`.
3. Klik **OK / Ya**.
4. **Ekspektasi Hasil**:
   * Rapat tersebut langsung terhapus dari daftar kartu.
   * Total rapat pada halaman profil berkurang.
   * Muncul toast sukses `"Notulensi rapat berhasil dihapus"`.
