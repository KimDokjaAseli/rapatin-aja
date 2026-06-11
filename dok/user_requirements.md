# 📋 Dokumen Kebutuhan Pengguna (User Requirements) - RapatIn

Dokumen ini menjelaskan semua fitur dan kebutuhan sistem yang harus dipenuhi oleh aplikasi **RapatIn** agar sesuai dengan kebutuhan pengguna.

---

## 👥 Aktor Sistem
* **User (Notulis / Pengembang)**: Orang yang menggunakan aplikasi untuk mencatat, melihat, mencari, mengedit, dan menghapus notulensi rapat, serta mengelola profil pribadinya.

---

## 🛠️ Kebutuhan Fungsional (Functional Requirements)

### 1. Sistem Autentikasi & Sesi
* **Pendaftaran Akun (Register)**:
  * User harus bisa membuat akun baru dengan mengisi Nama Lengkap, Email, dan Kata Sandi.
  * Sistem harus mengamankan kata sandi menggunakan enkripsi (hash bcrypt) sebelum disimpan ke database.
* **Masuk Aplikasi (Login)**:
  * User harus bisa masuk ke aplikasi menggunakan email dan kata sandi yang telah terdaftar.
* **Manajemen Sesi (Session)**:
  * Sistem harus menyimpan status login user di penyimpanan lokal (`localStorage` untuk web atau Secure Storage untuk mobile) agar user tidak perlu login ulang setiap kali membuka aplikasi.
* **Keluar Aplikasi (Logout)**:
  * User harus bisa keluar dari akun, menghapus data sesi di penyimpanan lokal, dan diarahkan kembali ke halaman login.

### 2. Pengelolaan Notulensi Rapat (CRUD)
* **Membuat Rapat Baru (Create)**:
  * User harus bisa membuat catatan rapat baru dengan mengisi: Judul Rapat, Tanggal (menggunakan kalender/date picker), Waktu, Lokasi, dan Deskripsi/Pembahasan.
  * Catatan rapat secara otomatis terhubung dengan ID user yang sedang login.
* **Melihat Daftar & Detail Rapat (Read)**:
  * **Dashboard/Home**: Menampilkan maksimal 4 rapat terbaru yang dibuat oleh user tersebut untuk akses cepat.
  * **Halaman Rapat**: Menampilkan daftar seluruh rapat yang dimiliki oleh user tersebut.
  * **Detail Rapat**: Menampilkan rincian lengkap rapat (tanggal, waktu, lokasi, pembahasan, keputusan, tindak lanjut).
* **Mencari Rapat**:
  * User harus bisa menyaring/mencari rapat secara langsung (*real-time*) berdasarkan Judul atau Lokasi rapat melalui kolom pencarian di bagian atas.
* **Mengubah Rapat (Update)**:
  * User harus bisa mengedit isi notulensi rapat. Saat form edit dibuka, data lama harus terisi secara otomatis.
* **Menghapus Rapat (Delete)**:
  * User harus bisa menghapus catatan rapat secara permanen.
  * Sistem harus menampilkan pesan konfirmasi terlebih dahulu sebelum proses penghapusan dilakukan.
  * Daftar rapat di layar harus ter-update secara otomatis setelah rapat dihapus tanpa perlu memuat ulang (*reload*) halaman.

### 3. Pengelolaan Profil Pengguna
* **Tampilan Profil**:
  * User harus bisa melihat detail profilnya sendiri yang berisi: Nama Lengkap, Jabatan/Posisi, Bio Singkat, Alamat Email, dan total statistik rapat yang pernah dibuat.
* **Edit Profil**:
  * User harus bisa mengubah informasi profilnya (Nama, Jabatan, Bio, Email). Perubahan ini harus langsung memperbarui database dan sesi lokal.
* **Foto Profil Otomatis (Avatar)**:
  * Sistem harus menampilkan gambar avatar inisial nama secara otomatis (menggunakan integrasi API gratis `ui-avatars.com`) jika user tidak mengunggah foto.

---

## 📈 Kebutuhan Non-Fungsional (Non-Functional Requirements)

### 1. Tampilan & Desain (UI/UX)
* **Tema Premium**: Aplikasi harus menggunakan skema warna Maroon modern dengan desain premium (seperti efek blur/glassmorphism dan transisi hover yang halus).
* **Responsif**: Tampilan aplikasi harus menyesuaikan ukuran layar secara otomatis agar nyaman digunakan di laptop, tablet, maupun HP.

### 2. Keamanan & Stabilitas
* **Keamanan Data**: Semua password wajib di-hash di backend menggunakan bcrypt.
* **CORS**: Backend harus dikonfigurasi dengan CORS agar frontend eksternal dapat melakukan request API dengan aman.

### 3. Kemudahan Penggunaan (Usability)
* **Notifikasi Instan (Toast/SnackBar)**: Aplikasi harus menampilkan pop-up singkat yang memberi tahu apakah suatu aksi (Login, Register, Simpan Rapat, Edit Profil, Hapus Rapat) berhasil atau gagal.
* **Indikator Loading**: Menampilkan spinner/loading status atau kartu skeleton agar user tahu aplikasi sedang memproses data saat mengambil data dari server.
