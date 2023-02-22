package main

import (
	"fmt"
)

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
	var array = []byte{}
	var result = false

	if len(s) > 1 {
		for i := 0; i < len(s); i++ {
			currentChar := s[i]
			if i > 0 {
				if array[len(array)-1] == '(' && currentChar == ')' {
					result = true
					array = func() []byte {
						var newArray = []byte{}
						for i := 0; i < len(array)-2; i++ {
							newArray = append(newArray, array[i])
						}
						return newArray
					}()
					continue
				} else if array[len(array)-1] == '{' && currentChar == '}' {
					result = true
					array = func() []byte {
						var newArray = []byte{}
						for i := 0; i < len(array)-2; i++ {
							newArray = append(newArray, array[i])
						}
						return newArray
					}()
					continue
				} else if array[len(array)-1] == '[' && currentChar == ']' {
					result = true
					array = func() []byte {
						var newArray = []byte{}
						for i := 0; i < len(array)-2; i++ {
							newArray = append(newArray, array[i])
						}
						return newArray
					}()
					continue
				} else if currentChar == '}' || currentChar == ')' || currentChar == ']' {
					result = false
					break
				}
			} else if currentChar == '}' || currentChar == ')' || currentChar == ']' {
				result = false
				break
			}
			array = append(array, currentChar)
		}

		if len(array) > 0 {
			result = false
		}
	}
	return result
}
