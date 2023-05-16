package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SebarapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luas)

	persegi := Persegi{Sisi: 4}
	luas = SebarapaLuas(persegi)
	fmt.Println("Luas Persegi: ", luas)

	asal := Persegi{Sisi: 4}
	luas = SebarapaLuas(asal)
	fmt.Println("Luas Persegi: ", luas)
}

func SebarapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
