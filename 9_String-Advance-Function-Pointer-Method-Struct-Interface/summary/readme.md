len() adalah sebuah fungsi built-in pada bahasa pemrograman Go yang digunakan untuk mengembalikan panjang sebuah string atau slice (urutan elemen dari tipe data yang sama).
==============contoh nya==========
package main

import "fmt"

func main() {
str := "Hello, world!"
strLen := len(str)
fmt.Println("Panjang string str adalah:", strLen)
}

compare
strings.Compare() adalah sebuah fungsi yang terdapat pada package strings
Fungsi ini digunakan untuk membandingkan dua string secara leksikografis, yaitu membandingkan karakter pertama pada kedua string dan kemudian karakter kedua dan seterusnya sampai menemukan perbedaan.
===============penggunaanya seperti ini ya...=================
package main

import "fmt"
import "strings"

func main() {
str1 := "mangga"
str2 := "pisang"

    result := strings.Compare(str1, str2)
    fmt.Println(result)

strings.Contains() yang terdapat pada package strings untuk mencari keberadaan sebuah substring dalam string.

Fungsi strings.Contains() mengambil dua argumen, yaitu string utama dan string yang akan dicari keberadaannya sebagai substring.
package main

import "fmt"
import "strings"

func main() {
str1 := "Hello, world!"
str2 := "world"

    if strings.Contains(str1, str2) {
        fmt.Printf("%s merupakan substring dari ", str2, str1)
    } else {
        fmt.Printf("%s bukan merupakan substring dari ", str2, str1)
    }

}
interface
Interface pada bahasa pemrograman Go adalah sebuah tipe data abstrak yang merepresentasikan kumpulan method atau perilaku yang harus diimplementasikan oleh sebuah tipe data konkret (struct) untuk memenuhi kontrak interface tersebut. Sebuah interface tidak memiliki implementasi atau nilai, hanya deklarasi method-method yang harus diimplementasikan.

method
Metho pada bahasa pemrograman Go adalah sebuah fungsi yang terasosiasi dengan sebuah tipe data (struct), sehingga fungsi tersebut dapat diakses sebagai sebuah property pada tipe data tersebut.

struct
untuk merepresentasikan objek-objek dalam program, dengan menyimpan data-data yang berkaitan dengan objek tersebut sebagai field pada struct. Field pada struct dapat diakses dan dimanipulasi secara langsung melalui operator titik (.) pada objek struct.

package main

import "fmt"

Deklarasi interface
type HitungLuas interface {
Luas() float64
}
type PersegiPanjang struct {
Panjang float64
Lebar float64
}
func (p PersegiPanjang) Luas() float64 {
return p.Panjang \* p.Lebar
}

Deklarasi tipe data konkret kedua yang mengimplementasikan interface HitungLuas
type Lingkaran struct {
JariJari float64
}

func (l Lingkaran) Luas() float64 {
return 3.14 _ l.JariJari _ l.JariJari
}

func main() {
Mendefinisikan array dari interface HitungLuas
hitungLuasArray := []HitungLuas{PersegiPanjang{4, 5}, Lingkaran{7

    Menghitung luas masing-masing objek pada array
    for _, objek := range hitungLuasArray {
        fmt.Printf(" objek.Luas())
    }

}
