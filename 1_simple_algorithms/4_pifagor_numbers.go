package main

import "fmt"

func main () {
	var(
		bord int = 20
	)	

	for i := 1; i < bord; i++ {
		for j := 1; j < bord; j++ {
			for k := 1; k < bord; k++ {
				if i*i + j*j == k*k && i < j {
					fmt.Printf("%d^2 * %d^2 = %d^2\n", i, j, k)
				}
			}
		}
	}
}