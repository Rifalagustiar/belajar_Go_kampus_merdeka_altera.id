package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	post := map[string]interface{}{
		"userid": 1,
		"title":  "Rifal ",
		"body":   "agustiar",
	}
	postJson, _ := json.Marshal(post)

	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postJson))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Respons status", response.Status)
}
