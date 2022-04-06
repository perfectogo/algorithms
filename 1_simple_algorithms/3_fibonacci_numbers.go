package main

import "fmt"

func fibo(ord int) int {
	var (
		f0    int = 0
		f1    int = 1
		f_ord int
	)
	for i := 0; i <= ord; i++ {
		f_ord = f0
		f0 = f1
		f1 = f0 + f_ord
	}
	return f_ord
}
func main() {
	var numOrder int = 10
	fmt.Println(fibo(numOrder))
}
