package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		cek := pow(x, n/2)
		return cek * cek
	}
	cek := pow(x, (n-1)/2)
	return x * cek * cek
}
func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
