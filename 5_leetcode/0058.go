package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("luffy is still joyboyf "))
}

func lengthOfLastWord(s string) int {
	var lastlen int
	var word string
	s += " "

	for _, run := range s {

		if run != ' ' {
			word += string(run)
			continue

		} else if word != "" {

			lastlen = len(word)
		}

		word = ""
	}

	return lastlen
}
