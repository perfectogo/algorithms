package main

import (
	"fmt"
)

func main () {
	var a, b, c int = 3, 7, 5

	if a > b && a > c {
		fmt.Printf("%d is greet\n", a)
		return
	} else if b > a && b > c {
		fmt.Printf("%d is greet\n", b)
		return
	}
	fmt.Printf("%d is greet\n", c)
}