package main

import "fmt"

func kelipatan(chanel chan int) {
	for i := 1; i <= 10; i++ {
		chanel <- i * 3
	}
	close(chanel)
}
func main() {
	chanel := make(chan int)
	go kelipatan(chanel)
	for lipat := range chanel {
		fmt.Println(lipat)
	}
}
