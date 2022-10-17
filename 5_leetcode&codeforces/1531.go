package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLengthOfOptimalCompression("aaccccabbaaaaa", 0))

}

func getLengthOfOptimalCompression(s string, k int) int {
	var (
		mapp  = make(map[string]int)
		lenth int
	)

	if len(s) == k {
		return 0
	}
	for _, run := range s {
		mapp[string(rune(run))] += 1
	}
	fmt.Println(mapp)
	return lenth
}

func sort(m map[string]int) {

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	fmt.Println(m)
}

// func getLengthOfOptimalCompression(s string, k int) int {
// 	var (
// 		chars  [128]int
// 		lenth  int
// 		hasGot bool
// 	)
// 	if len(s) == k {
// 		return 0
// 	}
// 	for _, run := range s {
// 		chars[run]++
// 	}

// 	for i, quantity := range chars {
// 		if quantity != k {
// 			if quantity == 1 {
// 				lenth += len(string(rune(i)))
// 				hasGot = true
// 				continue
// 			}

// 			if quantity > 1 {
// 				lenth += len(fmt.Sprint(quantity) + string(rune(i)))
// 				hasGot = true
// 			}
// 		}
// 	}

// 	if !hasGot {
// 		lenth = 1
// 	}

// 	return lenth
// }
