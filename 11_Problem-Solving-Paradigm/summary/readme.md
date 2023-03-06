Problem Solving Paradigm adalah pendekatan atau strategi dalam menyelesaikan suatu masalah pemrograman. Beberapa paradigma yang sering digunakan dalam pemrograman adalah Brute Force, Greedy, dan Dynamic Programming.
Brute Force
Brute Force adalah strategi penyelesaian masalah dengan cara mencoba semua kemungkinan solusi secara sistematis. Pada umumnya, metode Brute Force digunakan ketika masalah yang diberikan memiliki ukuran masukan yang kecil dan solusi yang optimal.
contoh
package main

import "fmt"

func sumOfTwo(arr []int, target int) bool {
n := len(arr)
for i := 0; i < n; i++ {
for j := i + 1; j < n; j++ {
if arr[i]+arr[j] == target {
return true
}
}
}
return false
}
Greedy
Greedy adalah strategi penyelesaian masalah dengan cara memilih solusi yang paling menguntungkan pada setiap tahap atau langkah. Pada umumnya, metode Greedy digunakan ketika masalah yang diberikan memiliki struktur optimal substructure.

package main

import (
"fmt"
"sort"
)

func maxSumSubarray(arr []int) int {
n := len(arr)
sort.Ints(arr)
sum := 0
for i := n - 1; i >= 0; i-- {
if arr[i] < 0 {
break
}
sum += arr[i]
}
return sum
}

func main() {
arr := []int{2, 4, -6, 5, -1, 3, 7, -8, 9}
fmt.Println(maxSumSubarray(arr))
}

Dynamic Programming
Dynamic Programming adalah strategi penyelesaian masalah dengan cara memecahkan masalah yang besar menjadi submasalah yang lebih kecil dan menyimpan solusi submasalah tersebut untuk digunakan kembali di kemudian hari. Pada umumnya, metode Dynamic Programming digunakan ketika masalah yang diberikan memiliki struktur overlapping subproblem.

package main

import "fmt"

func fibonacci(n int) int {
if n <= 1 {
return n
}
f := make([]int, n+1)
f[0] = 0
f[1] = 1
for i := 2; i <= n; i++ {
f[i] = f[i-1] + f[i-2]
}
return f[n]
}

func main() {
fmt.Println(fibonacci(10))
}
