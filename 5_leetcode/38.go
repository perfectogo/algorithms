package main

import "fmt"

func main() {
	fmt.Println(countAndSay(2))
}

var numbers = map[string]int{
	"1": 0, "2": 0, "3": 0,
	"4": 0, "5": 0, "6": 0,
	"7": 0, "8": 0, "9": 0,
}

// 1-1, 2-11, 3-21, 4-1211, 5-111221, 6-312211, 7-13112221, 8-1113213211 9-311312111321
func countAndSay(n int) string {
	var (
		i          int    = 2
		countedStr string = "1"
	)

	if n == 1 {
		countedStr = "1"
	}

	for i <= n {

	}
	fmt.Println(countedStr)
	return countedStr
}
