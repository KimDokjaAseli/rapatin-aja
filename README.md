# 📘 RapatIn – Dokumentasi Widget & Struktur Aplikasi

README ini menjelaskan struktur komponen utama dalam aplikasi **RapatIn** untuk mempermudah pengembangan, perawatan, dan kolaborasi.

---

## 1. Tema Global (AppTheme)

**Lokasi:** `lib/core/app_theme.dart`

File ini berfungsi sebagai pusat pengaturan tampilan aplikasi.

### Fungsi Utama
- Menyimpan warna utama, font, dan gaya tombol.
- Mengurangi duplikasi style pada setiap halaman.
- Perubahan warna aplikasi cukup dilakukan di file ini.

**Warna Utama:** Maroon (`0xFF800000`)

---

## 2. Widget Reusable

Widget di bawah ini digunakan berulang di banyak halaman.

### A. MeetingCard
**Lokasi:** `lib/widgets/meeting_card.dart`

**Fungsi:**
- Menampilkan ringkasan rapat dalam bentuk card.
- Memuat: Status rapat, judul, lokasi, dan tanggal.
- Digunakan pada halaman Home untuk daftar rapat.

---

### B. MeetingDetailSection
**Lokasi:** `lib/widgets/meeting_detail_section.dart`

**Fungsi:**
- Menampilkan bagian-bagian isi notulensi seperti:
  - Pembahasan
  - Keputusan
  - Tindak Lanjut
- Memiliki ikon dan warna tema yang dapat disesuaikan.

---

## 3. Form dan Validasi

### A. Form Login
**Lokasi:** `lib/pages/login_page.dart`

**Elemen:**
- `TextField` Username
- `TextField` Password

**Validasi:**
```dart
if (uname.isEmpty || pass.isEmpty) {
  ScaffoldMessenger.of(context).showSnackBar(
    const SnackBar(content: Text('Username dan Password harus diisi'))
  );
  return;
}
