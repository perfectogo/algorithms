package main

import (
	"fmt"
)

func BinarySearch(number int, numbers []int) int {
	var low, high int = 0, len(numbers)
	var middle int
	for low <= high {
		middle = (low + high) / 2
		item := numbers[middle]
		if item == number {
			return middle
		} else if item > number {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return 0
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 1000}
	number := 1000
	index := BinarySearch(number, numbers)
	fmt.Println(index)
}
