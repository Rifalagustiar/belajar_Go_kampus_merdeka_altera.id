package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func fetchProducts(url string, wg *sync.WaitGroup, ch chan<- Product) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		fmt.Printf("Error decoding data: %v\n", err)
		return
	}

	for _, product := range products {
		ch <- product
	}
}

func main() {
	url := "https://fakestoreapi.com/products"
	numWorkers := 10

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	ch := make(chan Product)

	for i := 0; i < numWorkers; i++ {
		go fetchProducts(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	productMap := make(map[string]int)

	for product := range ch {
		for _, char := range product.Title {
			productMap[string(char)]++
		}
	}

	for char, count := range productMap {
		fmt.Printf("%s: %d\n", char, count)
	}
}
