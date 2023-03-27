package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// hapus data per ID

	hapus := 1

	client := http.Client{}

	// request untuk hapus data dari API json

	request, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(hapus), nil)
	if err != nil {
		fmt.Println("gagal dalam membeuat request", err)
		return
	}
	// kirim request ke API
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Gagal dalam mengerim request", err)
		return
	}
	defer response.Body.Close()

	// kita tampilkan
	fmt.Println("status code:", response.StatusCode)
}
