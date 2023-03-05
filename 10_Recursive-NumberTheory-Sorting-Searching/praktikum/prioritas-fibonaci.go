package main

import "fmt"

func fibonaci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibonaci(number-1) + fibonaci(number-2)
	}
}
func main() {
	fmt.Println(fibonaci(0))
	fmt.Println(fibonaci(2))
	fmt.Println(fibonaci(9))
	fmt.Println(fibonaci(10))
	fmt.Println(fibonaci(12))
}
