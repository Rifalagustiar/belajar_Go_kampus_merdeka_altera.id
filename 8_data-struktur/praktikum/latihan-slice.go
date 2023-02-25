package main

import "fmt"

func mapping(silce []string) map[string]int {
	data := make(map[string]int)
	for _, v := range silce {
		data[v]++
	}
	return data
}
func main() {
	fmt.Println(mapping([]string{"rifal", "izam", "ridha", "lahu"}))
	fmt.Println(mapping([]string{"rifal", "izam", "ridha"}))
	fmt.Println(mapping([]string{}))
}
