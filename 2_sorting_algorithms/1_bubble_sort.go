package main

import "fmt" 

const(
	raise int = iota
	desc 
)

func bubbleSort(numbers []int, order int) []int {
	for i := 0; i < len(numbers) - 1; i++ {
		for j := 0; j < len(numbers) - i - 1; j++ {
			if order == raise && numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
			if order == desc && numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func main() {
	var numbers []int = []int{4, 2, 8, 16, 64, 32, 256, 128}
	fmt.Println(bubbleSort(numbers, raise))
	fmt.Println(bubbleSort(numbers, desc))
}

/*
output:
	[2 4 8 16 32 64 128 256]
	[256 128 64 32 16 8 4 2]
*/
