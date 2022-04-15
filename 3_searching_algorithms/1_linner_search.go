package main

import (
	"fmt"
)

type numberic interface{
	int | float64
}

func LinerSearch[T comparable](element T, elements []T) int {
	var steps int

	for i, e := range elements {
		steps++
		if element == e {
			return i
			break
		}
	}
	return 0
}

func main() {
	
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	number := 3
	index := LinerSearch[int](number, numbers)
	fmt.Println(index)

	words := []string{"stol", "stul", "oynak", "ko'zgu"}
	word := "oynak"
	index = LinerSearch[string](word, words)
	fmt.Println(index)
}
