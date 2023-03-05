package main

import (
	"fmt"
	"strconv"
)

func munculsekali(angka string) []int {

	nilai := make(map[rune]int)
	for _, v := range angka {
		nilai[v]++
	}
	var result []int
	for _, v := range angka {
		if nilai[v] == 1 {
			n, _ := strconv.Atoi(string(v))
			result = append(result, n)
		}
	}
	return result
}
func main() {
	fmt.Println(munculsekali("1234123"))
	fmt.Println(munculsekali("76523752"))
	fmt.Println(munculsekali("12345"))
	fmt.Println(munculsekali("1J75372645"))
	fmt.Println(munculsekali("71352675"))
}
