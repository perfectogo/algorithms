package main

import "fmt"

func fuctorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	} else {
		return number * fuctorial(number-1)
	}
}
func main() {
	var number int = 5
	fmt.Println(fuctorial(number))
}
