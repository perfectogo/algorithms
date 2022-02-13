package main

import (
	"fmt"
)

func main () {
	var (
		number, result int = 5, 1
	)

	i := number

	for i > 1 {
		result *= i
		i -= 1
	}

	fmt.Printf("%d! = %d", number, result)
}
