package main

import "fmt" 

func toFindMinMax(slc []int) (min_num, max_num int) {
	min_num = slc[0]
	max_num = slc[0]

	for _, item := range slc {
		if item < min_num {
			min_num = item
		}

		if item > max_num {
			max_num = item
		}
	}

	return min_num, max_num
}

func main() {
	var slc []int= []int{12, 34, 54, 70, 7,23, 56}

	fmt.Println(toFindMinMax(slc))
}