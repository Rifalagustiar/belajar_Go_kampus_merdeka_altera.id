package main

import "fmt"

func getminmax(number ...*int) (min int, max int) {
	min = *number[0]
	max = *number[0]
	for _, v := range number {
		if *v < min {
			min = *v
		}
		if *v > max {
			max = *v
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getminmax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min", min)
	fmt.Println("Nilai Min", max)
}
