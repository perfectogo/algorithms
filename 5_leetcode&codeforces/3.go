package main

import (
	"fmt"
)

func main() {
	word := "pwwkew"
	n := lengthOfLongestSubstring(word)
	fmt.Println(n)

}

func lengthOfLongestSubstring(s string) int {
	var (
		chars            [128]rune
		left, right, res int
	)

	for right < len(s) {
		r := s[right]
		chars[r]++

		for chars[r] > 1 {
			l := s[left]
			chars[l]--
			left++
		}
		if res < right-left+1 {
			res = right - left + 1
		}

		right++
	}

	return res
}

func lengthOfLongestSubstring_1(s string) int {
	var (
		s1  []rune
		idf bool
	)
	for _, ch := range s {
		for _, ch1 := range s1 {
			if ch == ch1 {
				idf = true
				break
			}
		}
		if !idf {
			fmt.Println(ch)
			s1 = append(s1, ch)
		}
		idf = false
	}
	return len(s1)
}

func lengthOfLongestSubstring_2(s string) int {
	var (
		n   = len(s)
		res int
	)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if res < j-i+1 {
					res = j - i + 1
				}
			}
		}
	}
	return res

}

// 0ce9448a-117e-46e8-9687-ec00b269b0f0 samarqaandadarvoza, minor, megaplanet, kingsmen, indenim
// etor
