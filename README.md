# Go Task Manager API

Sebuah project backend sederhana dengan **Go** yang menunjukkan:
- RESTful API menggunakan `net/http`
- Native SQL menggunakan `database/sql` (SQLite sebagai contoh)
- Worker background dengan **goroutine** untuk auto-archive task

---

## 🚀 Fitur
- Tambah task baru
- Ambil daftar task (yang belum diarsipkan)
- Tandai task selesai
- Background worker (goroutine) untuk auto-archive task selesai setiap 1 menit

---

## 📂 Struktur Project
