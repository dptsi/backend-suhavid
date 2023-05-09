package main //penamaan package berguna untuk class yang berkaitan. dan penamaan main untuk executable

import (
	"backend-suhavid/calculation"
	"fmt"
)

func main() {
	fmt.Println("Halo, suhavid hendra")

	result := calculation.Add(10, 9)

	fmt.Println(result)
}
