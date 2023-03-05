package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) EstimateDistance() float64 {
	// Menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar
	// yang sedang terisi, dengan asumsi efisiensi bahan bakar 1.5 L / mill
	return c.FuelIn * 1.5
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 20.0, // 20 liter bahan bakar tersedia
	}

	estimatedDistance := myCar.EstimateDistance() // menghitung perkiraan jarak yang bisa ditempuh
	fmt.Printf("Estimated distance: %.2f km\n", estimatedDistance)
}
