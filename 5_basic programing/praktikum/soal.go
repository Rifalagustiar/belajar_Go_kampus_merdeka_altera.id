package main

import "fmt"

func main() {
	grade := -5

	if grade > 80 {
		fmt.Println("nilai anda adalah A")
	} else if grade > 65 {
		fmt.Println("nilai anda adalah B")
	} else if grade > 50 {
		fmt.Println("nilai anda adalah C")
	} else if grade > 35 {
		fmt.Println("nilai anda adalah D")
	} else if grade > 0 {
		fmt.Println("nilai anda adalah E")
	} else {
		fmt.Println("nilai anda adalah tidak valid")
	}
}
