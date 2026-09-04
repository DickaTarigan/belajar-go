# CLI Todo App (Go)

Aplikasi pencatatan (to do list) berbasis CLI sederhana

## Fitur
- Menambahkan tugas dengan deteksi ID otomatis
- Melihat seluruh tugas yang sudah di input
- Update menandai tugas sudah selesai
- Hapus tugas
- Penyimpanan data lokal ('todos.json')

## Requirement
Pastikan komputer kamu sudah terpasang:
- [Go](https://go.dev/dl/) versi 1.20 atau yang lebih baru.

## Instalasi & Build

1. Clone repositori ini atau buka direktori proyek di terminal:
   ```bash
   cd cli-todo
   ```

2. Jalankan langsung menggunakan Go:
   ```bash
   go run main.go list
   ```

3. Atau kompilasi menjadi biner/executable:
   ```bash
   # Windows
   go build -o todo.exe main.go

   # Linux / macOS
   go build -o todo main.go
   ```

## Cara Penggunaan

### 1. Menambah Tugas
```bash
./todo add "Belajar goroutine dan channel"
```

### 2. Melihat Daftar Tugas
```bash
./todo list
```
Output:
```text
1. [ ] Belajar goroutine dan channel
```

### 3. Menandai Tugas Selesai
```bash
./todo done 1
```

### 4. Menghapus Tugas
```bash
./todo delete 1
```

## Teknologi yang Digunakan

- **Bahasa:** Go
- **Pustaka Standar:** `encoding/json`, `os`, `fmt`, `strconv`, `errors`