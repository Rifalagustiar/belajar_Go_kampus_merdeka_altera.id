package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	batu := make([]int, n)
	for i := range batu {
		batu[i] = math.MaxInt32
	}
	batu[0] = 0
	batu[1] = int(math.Abs(float64(jumps[1] - jumps[0])))
	for i := 2; i < n; i++ {
		batu[i] = int(math.Min(
			float64(batu[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))),
			float64(batu[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2])))),
		))
	}
	return batu[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))             //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 10, 50})) // 40
}
