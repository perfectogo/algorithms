package main

import "fmt"

func main() {
	var(
		num uint = 11
		i uint = 2
		idf bool = true
	)	

	for ; i < num; i++ {
		if num % i == 0 {
			idf = false
			break
		}
	}

	if idf {
		fmt.Printf("%d is prime number", num)
		return
	}

	fmt.Printf("%d is not prime number", num)
}