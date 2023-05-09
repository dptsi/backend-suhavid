package calculation

import "fmt"

func Add(number int, numberTwo int) int { //penamaan huruf awal kapital supaya bisa di akses package lain
	return number + numberTwo
}

// function
func result() {
	sentence := printMyResult("saya sedang")
	fmt.Println(sentence)
}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar Golang"
	return newSentence
}
