package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type post struct {
	UserID int    `json:"userID"`
	ID     int    `json:"ID"`
	Title  string `json:"Title"`
	Body   string `json:"body"`
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Printf("HTTP tidak bisa request %s\n", err)
	} else {
		defer response.Body.Close()

		var posts []post
		err := json.NewDecoder(response.Body).Decode(&posts)

		if err != nil {
			fmt.Printf("Error parsing the response body %s\n", err)
		} else {
			fmt.Println(posts)
		}
	}
}
