package main

import "fmt"

func dominant(n int, a [2]int) int {
	for i := 0; i < n; i++ {
		if a[i] == 1 {
			return 0
		}
	}
	return 1
}
func main() {
	var arr = [2]int{2, 2}
	fmt.Println(dominant(2, arr))
}
