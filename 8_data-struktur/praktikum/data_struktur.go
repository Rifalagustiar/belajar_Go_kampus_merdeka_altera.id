package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	data := map[string]bool{}
	datas := []string{}

	for _, i := range arrayA {
		if ya := data[i]; !ya {
			data[i] = true
			datas = append(datas, i)
		}
	}

	for _, i := range arrayB {
		if ya := data[i]; !ya {
			data[i] = true
			datas = append(datas, i)
		}
	}
	return datas
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jim", "Akuma"}, []string{"eddie", "stave", "gise"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jim", "Akuma"}, []string{"jin", "stave", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yotmitsu"}, []string{"devil jin", "yotmitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
