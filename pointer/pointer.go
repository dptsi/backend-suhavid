package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {
	number := 5
	fmt.Println("Alamat memory :", &number) // untuk memperoleh alamat memory menggunakan tanda "&"
	fmt.Println("Nilai awal :", number)

	change(&number, 100)

	fmt.Println("Nilai akhir :", number)
	fmt.Println("Alamat memory :", &number)

	//pointer struct sebagai parameter
	student := Student{1, "Jhon doe", 3.23}
	fmt.Println(student.Name)
	graduate(&student)

	fmt.Println(student.Name)
}

func change(old *int, new int) { // tipe data memory harus menggunakan * dibelakang tipe data/ direferencing
	*old = new
	fmt.Println("Alamat memory :", &old)
	fmt.Println("Di dalam function :", *old)
}

func graduate(student *Student) {
	student.Name = student.Name + " S.T"
}
