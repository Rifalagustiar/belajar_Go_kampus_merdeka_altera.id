package main

import "fmt"

func prima(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func primex(number int) int {
	count := 0
	i := 2
	for count < number {
		if prima(i) {
			count++
			if count == number {
				return i
			}
		}
		i++
	}
	return -1
}
func main() {
	fmt.Println(primex(1))
	fmt.Println(primex(5))
	fmt.Println(primex(8))
	fmt.Println(primex(9))
	fmt.Println(primex(10))
}
