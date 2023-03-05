package main

import "fmt"

func compare(a, b string) string {
	rifal := ""
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			rifal += string(a[i])
		}
	}
	return rifal
}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}
