package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("salom"))

}
func longestPalindrome(s string) (pl string) {

	for i := len(s); i <= 0; i-- {
		str, idf := re(s, i)
		if idf {
			pl = str
			break
		}
	}
	return pl
}

func re(str string, i int) (s string, idf bool) {
	return

}
