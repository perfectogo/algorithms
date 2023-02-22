package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(323))
}

func isPalindrome(x int) bool {
	var y, z = 0, x

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	for z > 0 {
		y *= 10
		y += z % 10
		z /= 10
	}

	if x == y {
		return true
	}

	return false
}
