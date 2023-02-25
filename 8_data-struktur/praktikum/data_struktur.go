package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	data1 := map[string]bool{}
	data2 := []string{}

	for _, i := range arrayA {
		if ya := data1[i]; !ya {
			data1[i] = true
		}
		data2 = append(data2, i)
	}

	for _, i := range arrayB {
		if ya := data1[i]; !ya {
			data1[i] = true
			data2 = append(data2, i)
		}
	}
	return data2
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jim", "Akuma"}, []string{"eddie", "stave", "gise"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jim", "Akuma"}, []string{"jin", "stave", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yotmitsu"}, []string{"devil jin", "yotmitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
