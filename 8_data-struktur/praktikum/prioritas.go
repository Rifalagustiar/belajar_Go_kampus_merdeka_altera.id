package main

import "fmt"

func pairsum(arr []int, target int) []int {

	chek := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		rifal := target - arr[i]
		if chek[rifal] {
			return []int{rifal, arr[i]}
		}
		chek[arr[i]] = true
	}
	return []int{}
}

func main() {
	fmt.Println(pairsum([]int{1, 2, 3, 4, 5, 6}, 6))
	fmt.Println(pairsum([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairsum([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairsum([]int{1, 4, 6, 8}, 10))
	fmt.Println(pairsum([]int{1, 5, 6, 7}, 6))
}
