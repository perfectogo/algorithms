package main

import "fmt"

var s int

func mult1(num1, num2 int) int {
	if num2 >= 1 {
		s = s + num1
		mult1(num1, num2-1)
		return s
	}
	return 0
}

func mult2(num1, num2 int) int {
	if num2 > 1 {
		return num1 + mult2(num1, num2-1)
	}
	return num1
}

func main() {
	fmt.Println(mult1(10, 5))
	fmt.Println(mult2(10, 5))
}
