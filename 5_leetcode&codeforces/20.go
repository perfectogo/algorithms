package main

import "fmt"

// true (){}[], {()[]}
var (
	s1 = "(){}[]"
	s2 = "({[]})"
	s3 = "({}[])}{"
	s4 = "({)}"
	s5 = "[({])}"
	ss = []string{s1, s2, s3, s4, s5}
)

func main() {
	for _, v := range ss {
		fmt.Println(isValid(v))
	}

}

func isValid(s string) bool {
	var (
	// i        int
	// statuses = map[string]int{
	// 	"(": 0,
	// 	"[": 0,
	// 	"{": 0,
	// 	")": 0,
	// 	"]": 0,
	// 	"}": 0,
	// }
	)
	return false
}

// func isValid(s string) bool {
// 	var (
// 		i        int
// 		statuses = map[string]int{
// 			"(": 0,
// 			"[": 0,
// 			"{": 0,
// 			")": 0,
// 			"]": 0,
// 			"}": 0,
// 		}
// 		brackets = map[string]string{
// 			"(": ")",
// 			"[": "]",
// 			"{": "}",
// 		}
// 		notBrackets = map[string]string{
// 			")": "(",
// 			"]": "[",
// 			"}": "{",
// 		}
// 	)

// 	for _, r := range s {
// 		statuses[string(r)]++
// 	}

// 	if !(statuses["("] == statuses[")"] &&
// 		statuses["{"] == statuses["}"] &&
// 		statuses["["] == statuses["]"]) {
// 		return false
// 	}
// 	// 0, 2; 3, 5; 6, 8
// 	for i < len(s) {
// 		if i+2 < len(s) {
// 			if brackets[string(s[i])] == string(s[i+2]) {
// 				return false
// 			}
// 		}
// 		if i+1 < len(s) {
// 			if notBrackets[string(s[i])] == string(s[i+1]) {
// 				fmt.Println(string(s[i]), string(s[i+1]))
// 				return false
// 			}
// 		}
// 		i++
// 	}
// 	return true
// }
