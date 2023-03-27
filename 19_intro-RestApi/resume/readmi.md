REST API (Representational State Transfer Application Programming Interface) merupakan sebuah teknologi yang digunakan untuk mengakses atau memanipulasi data dari sebuah server melalui protokol HTTP atau HTTPS. Pada bahasa pemrograman Go, implementasi REST API dapat dilakukan dengan menggunakan package net/http dan package encoding/json yang sudah disediakan oleh Go.

Berikut adalah beberapa hal yang perlu diketahui dalam mengimplementasikan REST API pada bahasa pemrograman Go:

Routing
Routing merupakan proses pengalihan permintaan (request) dari klien ke fungsi atau handler yang sesuai. Pada Go, package "net/http" digunakan untuk menghandle request HTTP. Fungsi http.HandleFunc() digunakan untuk menangani request yang masuk pada suatu path tertentu. Contoh penggunaan:

http.HandleFunc("/books", func(w http.ResponseWriter, r \*http.Request) {
// fungsi untuk menangani request HTTP pada path /books
})

Handling HTTP Methods
Pada REST API, HTTP Method yang umum digunakan adalah GET, POST, PUT, dan DELETE. Pada Go, kita dapat menangani request dengan metode HTTP tertentu dengan memanfaatkan fungsi yang disediakan oleh package "net/http". Contoh penggunaan:
HTTP GET Method
go
Copy code
http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
if r.Method == "GET" {
// mengembalikan semua data buku
}
})
HTTP POST Method
go
Copy code
http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
if r.Method == "POST" {
// menambahkan data buku baru
}
})
HTTP PUT Method
go
Copy code
http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
if r.Method == "PUT" {
// mengubah data buku berdasarkan id
}
})
HTTP DELETE Method
go
Copy code
http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
if r.Method == "DELETE" {
// menghapus data buku berdasarkan id
}
})
JSON Encoding dan Decoding
REST API pada umumnya menggunakan format data JSON untuk mengirim data antara server dan klien. Pada Go, package "encoding/json" digunakan untuk melakukan encoding dan decoding data JSON. Contoh penggunaan:
Encoding Struct ke JSON
go
Copy code
type Book struct {
ID int `json:"id"`
Title string `json:"title"`
Author string `json:"author"`
}

book := Book{
ID: 1,
Title: "The Lord of The Rings",
Author: "J.R.R. Tolkien",
}

bookJSON, \_ := json.Marshal(book)
Decoding JSON ke Struct
csharp
Copy code
type Book struct {
ID int `json:"id"`
Title string `json:"title"`
Author string `json:"author"`
}

var book Book
bookJSON := []byte(`{"id":1,"title":"The Lord of The Rings","author":"J.R.R. Tolkien"}`)

err := json.Unmarshal(bookJSON, &book)
Dalam implementasi REST API pada bahasa pemrograman Go, terdapat beberapa hal yang perlu diperhatikan seperti routing, handling HTTP methods, JSON encoding dan decoding. Selain itu, perlu juga diperhatikan mengenai penggunaan database dan konfigurasi server agar dapat berjalan dengan baik.
