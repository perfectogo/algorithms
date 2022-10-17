package main

import "fmt"

var RNum = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var i, sum int
	for i < len(s) {
		if i+1 < len(s) {
			if (i+1 < len(s)) &&
				(string(s[i]) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X")) ||
				(string(s[i]) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C")) ||
				(string(s[i]) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M")) {
				sum += RNum[string(s[i+1])] - RNum[string(s[i])]
				i += 2
				continue
			}
		}
		sum += RNum[string(s[i])]
		i++
	}
	return sum
}

/*
	//TIME
func romanToInt(s string) int {
    valueMap := map[rune]int {
        'M': 1000,
        'D': 500,
        'C': 100,
        'L': 50,
        'X': 10,
        'V': 5,
        'I': 1,
    }
	intermString := strings.Replace(s, "IV", "IIII", -1)
	intermString = strings.Replace(intermString, "IX", "VIIII", -1)
	intermString = strings.Replace(intermString, "XL", "XXXX", -1)
	intermString = strings.Replace(intermString, "XC", "LXXXX", -1)
	intermString = strings.Replace(intermString, "CD", "CCCC", -1)
	intermString = strings.Replace(intermString, "CM", "DCCCC", -1)
	mainCount := 0
	for _, romanChar := range intermString {
		mainCount += valueMap[romanChar]
	}
	return mainCount
}
*/
