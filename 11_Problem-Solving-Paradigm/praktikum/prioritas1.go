package main

import "fmt"

func generate(number int) [][]int {
	segitiga := make([][]int, number)
	for i := 0; i < number; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = segitiga[i-1][j-1] + segitiga[i-1][j]
		}
		segitiga[i] = row
	}
	return segitiga
}

func main() {
	number := 5
	segitigapaskal := generate(number)
	// menampilkan output segitiga Pascal
	for i := 0; i < number; i++ {
		fmt.Printf("%v\n", segitigapaskal[i])
	}
}
