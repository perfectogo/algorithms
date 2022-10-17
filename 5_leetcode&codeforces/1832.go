package main

import "fmt"

func main() {
	str := "jwtucoucmdfwxxqnxzkaxoglszmfrcvjoiunqqausaxxaaijyqdqgvdnqcaihwilqkpivenpnekioyqujrdrovqrlxovcucjqzjsxmllfgndfprctxvxwlzjtciqxgsxfwhmuzqvlksyuztoetyjugmswfjtawwaqmwyxmvo"
	fmt.Println(checkIfPangram(str))
}
func checkIfPangram(sentence string) (idf bool) {
	var (
		mapp = make(map[int]int)
	)
	idf = true

	if len(sentence) < 26 {
		idf = false
	}

	for i := 97; i <= 122; i++ {
		mapp[i] = 0
	}
	for _, s := range sentence {
		mapp[int(rune(s))]++
	}
	for _, v := range mapp {
		if v == 0 {
			idf = false
		}
	}
	return idf
}
