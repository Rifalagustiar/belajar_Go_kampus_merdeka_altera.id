package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var hitung, hitungs int
	for i := 0; i < len(matrix); i++ {
		hitung += matrix[i][i]
		// hitungs += matrix[i][len(matrix)[i]-1]
		hitungs += matrix[i][len(matrix)-i-1]
	}
	menghitung := aceh(hitung - hitungs)
	for _, tahu := range matrix {
		fmt.Println(tahu)
	}
	fmt.Println("selisih antara daigonal dan absolt adalah", menghitung)
}
func aceh(data int) int {
	if data < 0 {
		return -data
	}
	return data
}
