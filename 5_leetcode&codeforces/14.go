package main

import (
	"fmt"
	"strings"
)

var (
	strs  = []string{"flower", "flow", "f1light"}
	strs2 = []string{"dog", "racecar", "car"}
)

func main() {
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	minStr := strs[0]

	for _, str := range strs {
		if len(minStr) > len(str) {
			minStr = str
		}
	}
	for i := len(minStr); i >= 0; i-- {
		idf := true
		partStr := minStr[:i]
		for _, str := range strs {
			if !strings.Contains(str[:i], partStr) {
				idf = false
				break
			}
		}
		if idf {
			return partStr
		}
	}
	return ""
}
