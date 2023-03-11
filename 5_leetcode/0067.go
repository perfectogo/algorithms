package main

// 1+1 = 10, 10 + 10 = 100, 100 + 100 = 1000, 1001 + 1 = 1010
// 1010 + 1011 = 10101
import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("1111", "1111"))
}

func addBinary(a string, b string) string {
	var len_a, len_b int = len(a), len(b)
	var binNum string = " "
	var inMem = "0"

	var eq = func(s string, l1, l2 int) string {

		for i := 0; i < l1-l2; i++ {
			s = "0" + s
		}
		return s
	}

	if len_b < len_a {
		b = eq(b, len_a, len_b)
	} else {
		a = eq(a, len_b, len_a)
	}

	for i := len(a) - 1; i >= 0; i-- {

		if a[i] == '0' && b[i] == '0' {
			binNum = inMem + binNum
			inMem = "0"
			continue
		}

		if a[i] == '1' && b[i] == '1' {

			if inMem == "1" {
				binNum = inMem + binNum
			} else {
				binNum = "0" + binNum
			}
			inMem = "1"
			continue

		}

		{
			if inMem == "1" {
				binNum = "0" + binNum
			} else {
				binNum = "1" + binNum
			}

		}
	}

	if inMem == "1" {
		binNum = inMem + binNum
	}

	return binNum[:len(binNum)-1]
}
