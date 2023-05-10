package calculation

import (
	"errors"
	"fmt"
)

func studyCase() {
	scores := []int{10, 7, 3, 4}
	total := sum(scores)
	fmt.Println(total)

	result, err := calculate(10, 2, "*")
	if err != nil { // default error adalah nill
		fmt.Println("Terjadi kersalahan")
	}

	fmt.Println(result)

}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Unkown operation")
	}

	return result, errorResult
}
