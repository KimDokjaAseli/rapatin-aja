# 📘 Dokumentasi Widget & Validasi Aplikasi RapatIn

Dokumen ini menjelaskan struktur widget, file penting, dan validasi utama yang digunakan dalam aplikasi **RapatIn**. Tujuannya untuk memudahkan pengembangan, pemeliharaan, dan onboarding developer baru.

---

## 1. Core & Tampilan (Theme)

### `app_theme.dart`
**Isi Utama:**
- Warna utama aplikasi (Maroon).
- Gaya input (border bulat).
- Gaya tombol dan elemen visual lainnya.

**Kegunaan:**
- Menyatukan tampilan aplikasi dalam satu tempat.
- Menghindari duplikasi style pada setiap halaman.
- Perubahan tampilan dapat dilakukan dari satu file.

---

## 2. Halaman Autentikasi

### `login_page.dart`
**Komponen:**
- `TextField` Username
- `TextField` Password
- Tombol masuk

**Validasi:**
- Username dan Password wajib diisi.
- Jika kosong, tampil SnackBar: **"Username dan Password harus diisi"**.

---

### `register_page.dart`
**Komponen:**
- Nama Lengkap
- Username
- Password

**Validasi:**
- Semua kolom wajib diisi.
- SnackBar error: **"Semua field harus diisi"**.
- Password dianjurkan minimal 6 karakter.

---

## 3. Manajemen Rapat (Notulensi)

### `home_page.dart`
**Fitur Utama:**
- Menampilkan daftar rapat menggunakan `ListView` berisi `MeetingCard`.
- Pencarian rapat berdasarkan judul.

---

### `meeting_detail_page.dart`
**Komponen:**
- Menggunakan `CustomScrollView` dan `SliverAppBar`.
- Menampilkan detail lengkap: tanggal, lokasi, pembahasan, keputusan, tindak lanjut.

**Validasi & Aksi:**
- Terdapat dialog konfirmasi sebelum menghapus rapat.

---

### `meeting_form_page.dart`
Digunakan untuk **Tambah** maupun **Edit Rapat**.

**Komponen:**
- Form input multi-baris untuk notulensi.
- Pemilihan tanggal menggunakan `DatePicker`.

**Validasi:**
- **Judul Rapat** wajib diisi (pesan: *"Judul tidak boleh kosong"*).
- Kolom lain opsional untuk fleksibilitas.

---

## 4. Profil Pengguna

### `profile_page.dart`
**Komponen:**
- Foto profil (`CircleAvatar`)
- Nama dan jabatan
- Tombol Logout

---

### `edit_profile_page.dart`
**Fitur:**
- Form edit data diri.
- Pemilihan foto profil menggunakan `ImagePicker`.

**Validasi:**
- Nama lengkap tidak boleh dikosongkan.

---

## 5. Struktur Pendukung (Backend & Data)

### `main.dart`
- Titik awal aplikasi.
- Menetapkan halaman pertama (default: `LoginPage`).
- Menerapkan tema global dari `AppTheme`.

---

### `models/`
**Isi:** Struktur data seperti:
- `Meeting`
- `User`

Fungsinya sebagai blueprint data di seluruh aplikasi.

---

### `services/`
Berisi logika komunikasi dengan server (API):
- Mengambil data rapat
- Menambah data
- Mengedit data
- Menghapus data

Penempatan khusus ini menjaga UI tetap bersih dari logika backend.

---

## 6. Widget Kustom

### `meeting_card.dart`
**Fungsi:**
- Menampilkan ringkasan rapat di halaman Home.

**Logika Warna:**
- Status “Selesai”: ikon **Hijau**
- Status lainnya: ikon **Maroon**

---

### `meeting_detail_section.dart`
**Fungsi:**
- Menyusun bagian-bagian isi detail rapat (misalnya Keputusan Akhir).
- Menyediakan tampilan konsisten dengan ikon dan styling rapi.

---

## 7. Ringkasan Validasi Utama

1. **Pengecekan Kolom Kosong:** Hampir semua aksi simpan mengecek kolom utama terlebih dahulu.
2. **Feedback Visual:** Menggunakan `SnackBar` untuk pesan error atau sukses.
3. **Loading State:** Menampilkan `CircularProgressIndicator` saat proses berjalan.

---

## Kesimpulan Singkat
Struktur aplikasi RapatIn dibuat sederhana dan modular. Dengan memisahkan tema, widget, data model, dan layanan server, pengembangan aplikasi menjadi lebih mudah dikelola dan scalable.

