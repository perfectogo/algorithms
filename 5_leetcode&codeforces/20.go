package main

import "fmt"

// true (){}[], {()[]}
var (
	s1 = "(){}[]"
	s2 = "({[]})"
	s3 = "({}[])"
	s4 = "({)}"
	s5 = "({)}]"
)

func main() {

	fmt.Println(isValid(s1))
}

func isValid(s string) bool {
	var (
		statuses = map[string]bool{
			"(": false,
			"[": false,
			"{": false,
			")": true,
			"]": true,
			"}": true,
		}
	)
	fmt.Println(statuses)
	for _, r := range s {
		fmt.Println(r)
	}
	return false
}
