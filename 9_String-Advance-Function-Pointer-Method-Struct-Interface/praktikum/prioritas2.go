package main

import "fmt"

func caesar(offset int, input string) string {

	var belajar string
	for _, v := range input {
		if v == 0 {
			belajar += " "
		} else {
			data := int(v) - 97
			databaru := (data + offset) % 26
			ambil := rune(databaru + 97)
			belajar += string(ambil)
		}
	}
	return belajar
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
