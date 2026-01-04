🏆 Aplikasi Manajemen Turnamen E-Sports (CLI) – Golang

Aplikasi Command Line Interface (CLI) yang dibangun menggunakan bahasa pemrograman Go (Golang) untuk mengelola data tim e-sports dan jadwal pertandingan. 
Project ini menampilkan implementasi nyata konsep Struktur Data dan Algoritma seperti sorting dan searching tanpa menggunakan library eksternal.

------------------------------------------------------------------------

🚀 Gambaran Umum

Aplikasi ini memungkinkan pengguna untuk: - Mengelola data tim dan
klasemen turnamen - Mengatur jadwal pertandingan - Melakukan pencarian
dan pengurutan data menggunakan algoritma klasik - Menampilkan tim
dengan performa terbaik

Seluruh fitur dijalankan melalui terminal dengan alur interaksi yang
sederhana dan mudah dipahami.

------------------------------------------------------------------------

✨ Fitur Utama

Manajemen Tim: - Menampilkan klasemen tim berdasarkan poin - Menambahkan
tim baru - Memperbarui data tim (menang, kalah, poin) - Menghapus tim -
Mencari tim menggunakan Sequential Search dan Binary Search

Manajemen Jadwal Pertandingan: - Menambah jadwal pertandingan -
Menampilkan jadwal pertandingan - Memperbarui jadwal pertandingan -
Menghapus jadwal pertandingan

Analisis Performa: - Menampilkan tim dengan performa terbaik berdasarkan
poin tertinggi

------------------------------------------------------------------------

🧠 Algoritma dan Konsep yang Digunakan

Searching: - Sequential Search - Binary Search (data diurutkan terlebih
dahulu)

Sorting: - Selection Sort (berdasarkan poin) - Insertion Sort
(berdasarkan nama tim)

Konsep Pemrograman: - Struct - Slice - Modular function - Validasi
input - Menu berbasis CLI

------------------------------------------------------------------------

🧱 Struktur Data

Struct Tim: type Tim struct { NamaTim string Menang int Kalah int Poin
int }

Struct JadwalTanding: type JadwalTanding struct { TimSatu string TimDua
string WaktuTanding string }

------------------------------------------------------------------------

▶️ Cara Menjalankan Program

Prasyarat: - Go versi 1.20 atau lebih baru

Langkah: 1. Clone repository 2. Masuk ke folder project 3. Jalankan
perintah: go run main.go

------------------------------------------------------------------------

🎯 Tujuan Pengembangan

-   Menerapkan konsep Struktur Data dan Algoritma
-   Melatih pemrograman Golang berbasis CLI
-   Menjadi bagian dari portfolio pengembangan software

------------------------------------------------------------------------

👨‍💻 Author

Nadiv
Mahasiswa Teknik Informatika
Fokus: Backend Development & Software Engineering

------------------------------------------------------------------------

Catatan: - Data masih bersifat in-memory - Aplikasi dijalankan melalui
terminal - Dapat dikembangkan lebih lanjut ke versi database atau web
