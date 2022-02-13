package main

import (
	"fmt"
)

func main () {
	var (
		numbers []int = []int {12, 24, 45, 57, 68, 79, 80, 99, 121}
		low, high, number int = 0, len(numbers) - 1, 80
		middle, index, steps int
	)

	for /*i := low; i < high; i++ && */low <= high {
		steps++
		middle = (low + high) / 2
		fmt.Println(middle)

		item := numbers[middle]
		if item == number {
			index = middle
			break
		} else if item < number {
			high = middle - 1
		} else {
			low = middle + 1
		}

	}

	fmt.Println("index", index)
	fmt.Println("item", numbers[index])
	fmt.Println("number" ,number)
	fmt.Println("steps", steps)
}