package main

import "fmt"

type Kendaraan struct {
	TotalRoda       int // mengikuti konvensi penamaan untuk TotalRoda dengan huruf besar
	KecepatanPerJam int // mengubah penamaan field agar sesuai dengan gaya bahasa Go
}

type Mobil struct {
	Kendaraan // menggunakan embedding untuk menyertakan struct Kendaraan ke dalam struct Mobil
}

func (m *Mobil) Berjalan() { // mengubah nama method agar sesuai dengan gaya bahasa Go
	m.tambahKecepatan(10)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) { // mengubah nama method agar sesuai dengan gaya bahasa Go
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{} // mengubah penulisan variabel agar sesuai dengan gaya bahasa Go
	mobilCepat.Berjalan()
	fmt.Printf("Mobil cepat: %dkm/jam\n", mobilCepat.KecepatanPerJam)
	mobilCepat.Berjalan()
	fmt.Printf("Mobil cepat: %dkm/jam\n", mobilCepat.KecepatanPerJam)
	mobilCepat.Berjalan()
	fmt.Printf("Mobil cepat: %dkm/jam\n", mobilCepat.KecepatanPerJam)

	mobilLambat := Mobil{} // mengubah penulisan variabel agar sesuai dengan gaya bahasa Go
	mobilLambat.Berjalan()
	fmt.Printf("Mobil lambat: %dkm/jam\n", mobilLambat.KecepatanPerJam)
}
