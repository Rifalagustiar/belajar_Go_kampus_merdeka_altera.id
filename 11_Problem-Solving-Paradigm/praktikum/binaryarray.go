package main

import "fmt"

func binaryArray(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = nilai(i)
	}
	return ans
}

func nilai(n int) int {
	Nilai := 0
	for n > 0 {
		if n%2 == 1 {
			Nilai++
		}
		n /= 2
	}
	return Nilai
}

func main() {
	arr := binaryArray(2)
	fmt.Println(arr)

	arr2 := binaryArray(3)
	fmt.Println(arr2)
}
