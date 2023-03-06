package main

import "fmt"

func toRoman(num int) string {
	var angka = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var simbol_romawi = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var hasil string

	for i := 0; i < len(angka); i++ {
		for num >= angka[i] {
			hasil += simbol_romawi[i]
			num -= angka[i]
		}
	}

	return hasil
}

func main() {
	fmt.Println(toRoman(4))
	fmt.Println(toRoman(9))
	fmt.Println(toRoman(23))
	fmt.Println(toRoman(2021))
	fmt.Println(toRoman(1646))
}
