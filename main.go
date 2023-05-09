package main //penamaan package berguna untuk class yang berkaitan. dan penamaan main untuk executable

import (
	"backend-suhavid/calculation"
	"fmt"
)

func main() {
	fmt.Println("Halo, suhavid hendra")

	result := calculation.Add(10, 9) // tanda := untuk membuat variabel baru tanpa menyertakan tipe data

	fmt.Println(result)

	//	If conditional
	score := 30
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)

	// Switch conditional
	number := 5

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Default")
	}

	//Perulangan for

	//for biasa
	for i := 1; i <= 100; i++ {
		fmt.Println("mencoba perulangan Go :", i)
	}

	// while
	i := 1
	for i <= 100 {
		fmt.Println("mencoba perulangan Go :", i)
		i++
	}

	// foreach
	title := "Golang the best language"
	for index, letter := range title {
		fmt.Println("index :", index, " letter ", string(letter))
	}

	// jika index tidak digunakan
	for _, letter := range title {
		fmt.Println(" letter :", string(letter))
	}

	// for dengan kondisi
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, " letter ", string(letter))
		}
	}

	// for dengan kondisi while
	for index, letter := range title {
		letterString := string(letter)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, " letter ", string(letter))
		}
	}

	// array for
	languages := [...]string{
		"Ruby",
		"C++",
		"PHP",
		"GO",
	}

	for index, lang := range languages {
		fmt.Println("Index : ", index, " language : ", lang)
	}

	length := len(languages)
	fmt.Println(length)

	// slice
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "PlayStation5")
	gamingConsole = append(gamingConsole, "PlayStation2")
	fmt.Println(gamingConsole)

	gaming := []string{"playstation2", "xbox"}
	fmt.Println(gaming)

	for _, console := range gaming {
		fmt.Println(console)
	}

	// map
	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"javascript": "hype",
	}

	delete(myMap, "ruby") // delete salah satu key dalam map

	for key, value := range myMap {
		fmt.Println("key : ", key, "  value: ", value)
	}

	// hitung rata-rata
	scores := [...]int{100, 80, 73, 82, 32, 63}
	var total int

	for _, score := range scores {
		total = total + score
	}

	jumlah := len(scores)
	rataRata := float64(total) / float64(jumlah) //float64 untuk menjadikan tipedata int ke float

	fmt.Println(rataRata)

}
