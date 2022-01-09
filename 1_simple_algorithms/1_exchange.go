package main

import "fmt"

func toExchange1 (num1, num2 int) (int, int) {
	holder := num1
	num1 = num2
	num2 = holder

	return num1, num2
}

func toExchange2 (num1, num2 int) (int, int) {
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	return num1, num2
}

// golang syntax
func toExchange3 (num1, num2 int) (int, int) {
	num1, num2 = num2, num1

	return num1, num2
}

func main() {
	var num1, num2 int = 5, 7

	fmt.Println(num1, num2) // 5, 7

	fmt.Println(toExchange1(num1, num2)) // 7, 5
	fmt.Println(toExchange2(num1, num2)) // 7, 5
	fmt.Println(toExchange3(num1, num2)) // 7, 5
}
