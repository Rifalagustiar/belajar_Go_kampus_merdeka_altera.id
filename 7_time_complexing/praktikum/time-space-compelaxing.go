package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		check := pow(x, n/2)
		return check * check
	}
	check := pow(x, (n-1)/2)
	return x * check * check
}
func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
