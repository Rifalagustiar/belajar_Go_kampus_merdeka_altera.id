package main

import (
	"fmt"
	"sync"
)

func hitungFrekuensi(huruf int32, Text string, chanel chan int, Tunggu *sync.WaitGroup) {
	defer Tunggu.Done()
	Hitung := 0
	for _, c := range Text {
		if c == huruf {
			Hitung++
		}
	}
	chanel <- Hitung
}

func main() {
	Text := "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	// Konversi Text ke dalam bentuk rune slice
	teks2 := []int32(Text)
	// Buat channel untuk menyimpan hasil perhitungan frekuensi
	chanel := make(chan int, len(teks2))
	// Buat WaitGroup untuk sinkronisasi goroutine
	var Tunggu sync.WaitGroup
	// Hitung frekuensi untuk setiap huruf dalam Text secara bersamaan
	for _, huruf := range teks2 {
		Tunggu.Add(1)
		go hitungFrekuensi(huruf, Text, chanel, &Tunggu)
	}
	// Tunggu hingga semua goroutine selesai
	Tunggu.Wait()
	// Tutup channel setelah semua goroutine selesai
	close(chanel)
	// Tampilkan hasil perhitungan
	for _, Hitung := range teks2 {
		if Hitung != ' ' {
			fmt.Printf("%c: %d\n", Hitung, <-chanel)
		}
	}
}
