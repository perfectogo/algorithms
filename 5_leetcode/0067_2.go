package main

// 1+1 = 10, 10 + 10 = 100, 100 + 100 = 1000, 1001 + 1 = 1010
// 1010 + 1011 = 10101
import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	var binNum string = " "
	var inMem = "0"
	var lenth int = len(b) - 1

	if len(a) < len(b) {
		lenth = len(a) - 1
	}

	for i := lenth; i >= 0; i-- {

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
