tipe data dan variable
import digunakan untuk meng-import atau memasukan package lain ke dalam file program, agar isi dari package yang di-import bisa dimanfaatkan.
Package fmt merupakan salah satu package bawaan yang disediakan oleh Go, isinya banyak fungsi untuk keperluan I/O yang berhubungan dengan text.
Berikut adalah skema penulisan keyword impor

proyek harus ada file program yang di dalamnya berisi sebuah fungsi bernama main(). Fungsi tersebut harus berada di file yang package-nya bernama <main>
fmt.Println() digunakan untuk memunculkan text ke layar (pada konteks ini, terminal 
fmt.Println() berada dalam package fmt, maka untuk menggunakannya perlu package tersebut untuk di-import terlebih dahulu.
  variable
variabel dideklarasikan dan diisi nilainya. (Keyword var) atau (:=) digunakan untuk membuat variabel baru.
  tipe data
tipe data seperti contoh int dan string
  
 konstanta
 Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya.
  
package main
import "fmt"

func main() {
	// OPERAN DAN OPETAOR
	// x := 1 + 2
	// fmt.Println(x)
// EXPRESION
	// a := 3
	// b := 5
	// x := a + b
	// fmt.Println(x)
// LUAS SEGITIGA
	// p := 20
	// l := 10
	// hasil := (p * l)/2
	// fmt.Println(hasil)
// OPERATION STRING
	rifal := "hello" + " " + "Agustiar"
	fmt.Println(rifal)
}
  variable
  package main

import "fmt"

func main() {
			// tipe data number
				// fmt.Println("angka", 1)
				// fmt.Println("angka", 2)
				// fmt.Println("angka", 3.5)

			// bolean true/ false
				// fmt.Println("benar", true)
				// fmt.Println("salah", false)

			// tipe data srting
				// fmt.Println("Rifal ")
				// fmt.Println("Agustiar")

			// variabel pada golang (tidak wajib)
				// var nama string 
				// nama = "rifal agustiar"
				// fmt.Printf(nama)
				
			// variabel wajib
				// nama := "Rifal agustiar"
				// fmt.Println(nama)

			// multiple variable
				// var (
				// 	nama_depan = "Rifal"
				// 	nama_belakang = "Agustiar"
				// )
				// fmt.Println(nama_depan)
				// fmt.Println(nama_belakang)
