package main

import (
	"fmt"
	"time"
)

func waktu(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * x)
		time.Sleep(3 * time.Second)
	}
}
func main() {
	go waktu(4)
	go waktu(5)
	time.Sleep(10 * time.Second)
}
