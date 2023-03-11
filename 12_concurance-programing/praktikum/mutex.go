package main

import (
	"fmt"
	"sync"
)

func main() {
	var penerapan sync.WaitGroup
	var mutex sync.Mutex
	counter := 0

	for i := 0; i < 10; i++ {
		penerapan.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
			fmt.Println("nilai:", counter)
			penerapan.Done()
		}()
	}

	penerapan.Wait()
	fmt.Println("total nilai:", counter)
}
