package main

import "fmt"

func mapping(silce []string) map[string]int {
	nilai := make(map[string]int)
	for _, v := range silce {
		nilai[v]++
	}
	return nilai
}
func main() {
	fmt.Println(mapping([]string{"rifal", "izam", "ridha", "lahu"}))
	fmt.Println(mapping([]string{"rifal", "izam", "ridha"}))
	fmt.Println(mapping([]string{}))
}
