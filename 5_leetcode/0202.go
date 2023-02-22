/*
Input: n = 19
Output: true
Explanation:
	1^2 + 9^2 = 82
	8^2 + 2^2 = 68
	6^2 + 8^2 = 100
	1^2 + 0^2 + 02 = 1
*/

package main

import "fmt"

func main() {
	fmt.Println(isHappy(7))
}

func isHappy(n int) bool {

	if isHappyRec(n) == 1 {
		return true
	}

	return false
}

func isHappyRec(n int) int {
	var (
		sum int
	)

	for n >= 1 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	if sum < 7 {
		return sum
	}

	return isHappyRec(sum)
}
