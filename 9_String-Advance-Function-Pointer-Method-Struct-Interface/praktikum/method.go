package main

import "fmt"

type Car struct {
	Type   string
	minyak float64
}

func (c *Car) EstimateDistance() float64 {

	return c.minyak * 1.5
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		minyak: 20.0,
	}

	estimatedDistance := myCar.EstimateDistance()
	fmt.Println("Estimated distance: km", estimatedDistance)
}
