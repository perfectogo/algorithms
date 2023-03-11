package main

import "fmt"

// 1 -> 1
// 2 -> 1-1, 2
// 3 -> 1-2, 2-1, 1-1-1
// 4 -> 1-1-1-1, 1-1-2, 1-2-1, 2-1-1, 2-2
// 5 -> 1-1-1-1-1, 1-1-1-2, 1-1-2-1, 1-2-1-1, 2-1-1-1, 1-2-2, 2-1-2, 2-2-1,
// 1, 2, 3, 5, 8, 13

func main() {

	fmt.Println(climbStairs(4))

}

func climbStairs(n int) int {
	var a, b = 1, 1

	for n > 0 {
		b, a = a, a+b
		n--
	}

	return b
}
