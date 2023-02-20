package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		//melakukan pengulangan bintang(*) sampai 30

		for j := 5; j >= i; j-- {
			fmt.Printf(" ")
		}

		for r := 1; r <= i; r++ {
			fmt.Printf("*")

		}
		for r := (i - 1); r >= 1; r-- {
			fmt.Printf("*")

		}
		fmt.Printf("\n")
	}
}
