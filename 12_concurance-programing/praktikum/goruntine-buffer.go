package main

import "fmt"

func cetak(chanel chan int) {
	for i := 1; i <= 10; i++ {
		chanel <- i * 3
	}
	close(chanel)
}

func main() {
	chanel := make(chan int, 10)
	go cetak(chanel)
	for kel3 := range chanel {
		fmt.Println(kel3)
	}
}
