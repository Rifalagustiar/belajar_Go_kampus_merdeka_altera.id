package main

import "fmt"

type pair struct {
	name  string
	count int
}

func mostapeartitem(barangs []string) []pair {
	nilai := make(map[string]int)

	for _, barang := range barangs {
		nilai[barang]++
	}
	var ambil []pair
	for name, hasil := range nilai {
		ambil = append(ambil, pair{name, hasil})
	}
	return ambil
}

func main() {
	fmt.Println(mostapeartitem([]string{"js", "js", "golang", "ruby", "js", "js"}))
	fmt.Println(mostapeartitem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostapeartitem([]string{"football", "basketball", "tenis"}))
}
