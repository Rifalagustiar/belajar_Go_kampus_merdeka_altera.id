package main

import "fmt"

func Palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	var str string
	fmt.Scan(&str)
	result := Palindrome(str)
	if result == true {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}

}
