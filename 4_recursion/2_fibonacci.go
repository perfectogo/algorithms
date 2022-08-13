package main

import "fmt"

func fibo(ord int) int {
	if ord <= 2 {
		return 1
	} else {
		return fibo(ord-2) + fibo(ord-1)
	}
}

func main() {
	var ord int = 5
	fmt.Println((ord))
}
// check ssh
