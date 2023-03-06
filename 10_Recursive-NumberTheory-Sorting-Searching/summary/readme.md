Recursive adalah teknik pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan sebuah permasalahan. Fungsi rekursif sering digunakan dalam pemrograman untuk menyelesaikan permasalahan yang berulang, seperti mencari faktorial atau nilai Fibonaci.
contoh penggunaan

package main

import "fmt"

func factorial(n int) int {
if n <= 1 {
return 1
} else {
return n \* factorial(n-1)
}
}

func main() {
fmt.Println(factorial(5))
}

Number Theory adalah cabang matematika yang mempelajari sifat-sifat bilangan bulat. Dalam pemrograman, konsep-konsep dari Number Theory sering digunakan untuk menyelesaikan permasalahan yang melibatkan operasi bilangan bulat, seperti mencari bilangan prima atau faktor suatu bilangan.

contoh penggunaan

package main

import "fmt"

func contoh(n int) bool {
if n <= 1 {
return false
}
for i := 2; i <= n/2; i++ {
if n%i == 0 {
return false
}
}
return true
}

func main() {
for i := 1; i <= 100; i++ {
if contoh(i) {
fmt.Print(i, " ")
}
}
}

Sorting adalah proses mengurutkan elemen-elemen dalam sebuah kumpulan data. Go menyediakan beberapa algoritma sorting yang siap pakai, seperti bubble sort, quick sort, dan merge sort. Sorting sering digunakan dalam pemrograman untuk mempermudah pengelolaan data dan pencarian data tertentu.

contoh pengunaan
package main

import "fmt"

func penggunaansorting(arr []int) {
n := len(arr)
for i := 0; i < n; i++ {
for j := 0; j < n-i-1; j++ {
if arr[j] > arr[j+1] {
arr[j], arr[j+1] = arr[j+1], arr[j]
}
}
}
}

func main() {
arr := []int{64, 34, 25, 12, 22, 11, 90}
penggunaansorting(arr)
fmt.Println("", arr)
}

Searching adalah proses mencari nilai tertentu dalam kumpulan data. Go menyediakan beberapa algoritma searching yang siap pakai, seperti binary search dan linear search. Searching sering digunakan dalam pemrograman untuk mencari data tertentu dalam database atau dalam kumpulan data yang besar.

package main

import "fmt"

func searching(arr []int, x int) int {
panjang := 0
lebar := len(arr) - 1

    for panjang <= lebar {
        contoh := (panjang + lebar) / 2

        if arr[contoh] < x {
            panjang = contoh + 1
        } else if arr[contoh] > x {
            lebar = contoh - 1
        } else {
            return contoh
        }
    }

    return -1

}

func main() {
arr := []int{11, 12, 22, 25, 34, 64, 90}
x := 25
result := searching(arr, x)
if result == -1 {
fmt.Printf(" x)
} else {
fmt.Printf( x, result)
}
}
