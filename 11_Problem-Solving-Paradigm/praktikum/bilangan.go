package main

import "fmt"

func SimpleEquetions(a, b, c int) {
	for x := -10000; x <= 10000; x++ {
		for y := -10000; y <= 10000; y++ {
			if x == y {
				continue
			}
			z := a - x - y
			if x*y*z == b && x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("no solution")
}
func main() {
	SimpleEquetions(1, 2, 3)
	SimpleEquetions(6, 6, 14)
}
