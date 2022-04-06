package main

import "fmt"

func div(num1, num2 int) int {
	if num1 > num2 {
		return 1 + div(num1-num2, num2)
	}
	return 1
}

func main() {
	fmt.Println(div(125, 25))
}

