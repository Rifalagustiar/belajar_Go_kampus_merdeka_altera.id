Prinsip Single Responsibility Principle (SRP)

Prinsip SRP pada arsitektur clean/hexagonal membagi aplikasi menjadi beberapa lapisan yang masing-masing bertanggung jawab atas satu tugas khusus. Setiap lapisan memainkan peran tertentu dalam arsitektur. Sebagai contoh, lapisan Repository bertanggung jawab untuk mengelola data, sementara lapisan Use Case bertanggung jawab untuk mengatur bisnis logic aplikasi. Prinsip SRP memungkinkan setiap lapisan untuk berubah dan diperbarui secara independen tanpa mempengaruhi lapisan lain.

Prinsip Dependency Inversion Principle (DIP)

Prinsip DIP pada arsitektur clean/hexagonal memungkinkan penggunaan teknologi atau library lain yang kompatibel dengan aplikasi tanpa merusak konsistensi arsitektur. Prinsip DIP menciptakan lapisan antarmuka (interface) antara komponen aplikasi yang memungkinkan komunikasi antara lapisan. Penggunaan interface memungkinkan untuk menghindari ketergantungan langsung antara komponen aplikasi. Sebagai contoh, lapisan Repository dan Use Case berkomunikasi melalui interface, yang memungkinkan penggunaan database berbeda tanpa mengubah kode di lapisan Use Case.

Struktur Aplikasi

Arsitektur clean/hexagonal pada Go-lang terdiri dari beberapa package dengan fungsi dan tanggung jawab masing-masing. Beberapa package yang umum digunakan pada arsitektur clean/hexagonal meliputi:

domain: berisi definisi model data atau struct yang digunakan pada aplikasi
repository: berisi logika untuk mengakses database atau penyimpanan data lainnya
usecase: berisi logika bisnis atau use case yang mengatur interaksi antara lapisan aplikasi
delivery: berisi kode untuk menangani interaksi antara aplikasi dengan dunia luar seperti HTTP, gRPC atau protokol lainnya
Struktur aplikasi pada arsitektur clean/hexagonal pada Go-lang memungkinkan pengembangan yang modular dan memudahkan dalam melakukan testing aplikasi secara terpisah pada setiap lapisan.
