package main

import (
	"fmt"
)

func main () {
	var index, steps int
	words := []string {"stol", "stul", "oynak", "ko'zgu"}
	word := "stul"

	for i, w := range words {
		steps++
		if word == w {
			index = i
			break
		}
	}

	if words[index] == word {
		fmt.Println("mission completed, steps:", steps)
		return
	}

	fmt.Println("mission not completed, steps:", steps)
}
