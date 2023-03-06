package main

import "fmt"

type Student struct {
	nama  string
	score []int
}

// cari nilai total maximum
func (siswa *Student) totalscore() int {

	total := 0
	for _, score := range siswa.score {
		total += score
	}
	return int(total) / int(len(siswa.score))
}

// cari total nilai minimum
func (siswa *Student) minimum() int {

	minimum := siswa.score[0]
	for _, score := range siswa.score {
		if score < minimum {
			minimum = score
		}
	}
	return minimum
}

// baru kita cara nilai maximum
func (siswa *Student) maximum() int {

	maximum := siswa.score[0]
	for _, score := range siswa.score {
		if score > maximum {
			maximum = score
		}
	}
	return maximum
}

// menginput nilai masing masing siswa
func main() {

	siswa1 := Student{nama: "Rizki", score: []int{80}}
	siswa2 := Student{nama: "Kobar", score: []int{75}}
	siswa3 := Student{nama: "Ismail", score: []int{70}}
	siswa4 := Student{nama: "Umam", score: []int{75}}
	siswa5 := Student{nama: "Yopan", score: []int{60}}

	students := []Student{siswa1, siswa2, siswa3, siswa4, siswa5}
	var total int
	var minimum, maximum int
	for _, student := range students {
		total += student.totalscore()
		if student.minimum() < minimum || minimum == 0 {
			minimum = student.minimum()
		}
		if student.maximum() > maximum {
			maximum = student.maximum()
		}
	}
	totalscore := total / int(len(students))

	// Menampilkan hasil aau output
	fmt.Println("Average score: ", totalscore)
	fmt.Println("Min score of student: Yopan", minimum)
	fmt.Println("Mix score of student: Yopan", maximum)
}
