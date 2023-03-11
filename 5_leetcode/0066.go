package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{9, 8, 9}))
}

func plusOne(digits []int) []int {

	var lenth, m int = len(digits), 1

	if digits[lenth-1] < 9 {
		digits[lenth-1] += 1
		return digits
	}

	if !func() bool {
		for i := lenth - 1; i >= 0; i-- {
			fmt.Println(digits[i])
			if digits[i] < 9 {
				return true
			}
		}
		return false
	}() {
		digits = append(digits, 0)
		lenth += 1
		m += 1
	}

	for i := lenth - m; i >= 0; i-- {
		fmt.Println(digits)
		if digits[i]+1 == 10 && i-1 >= 0 {

			digits[i] = 0

		} else if digits[i]+1 == 10 && i-1 < 0 {
			digits[i] = 1
			break
		} else {
			digits[i] += 1
			break
		}
	}

	return digits
}

/*
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] + 1 < 10 {
            digits[i] += 1
            return digits
        } else {
            digits[i] = 0
        }
    }
    if digits[0] == 0 {
        digits = append([]int{1}, digits...)
    }
    return digits
}
*/
