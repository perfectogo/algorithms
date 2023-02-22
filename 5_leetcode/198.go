package main

import (
	"fmt"
)

func main() {
	numsMap := make(map[int]bool)
	fmt.Println(numsMap[23])
	nums := []int{4, 2, 4, 4}

	ans := rob(nums)

	fmt.Println("THIS IS ans: ", ans)
}

func rob(nums []int) int {

	var (
		sum, max, indexBefore, indexAfter, maxIndex int
	)

	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numsMap[i] = nums[i]
	}

	for i := 0; i < len(numsMap); i++ {
		fmt.Println("iter: ", i)

		for j := 0; j < len(numsMap); j++ {
			if numsMap[j] > max {
				max = numsMap[j]
				maxIndex = j
				indexBefore = j - 1
				indexAfter = j + 1
			}
		}
		sum += max
		fmt.Println("this is max: ", max)
		max = 0

		numsMap[maxIndex] = 0
		numsMap[indexBefore] = 0
		numsMap[indexAfter] = 0

		fmt.Println("this is sum: ", sum)
		fmt.Println("this is indexBefore: ", indexBefore)
		fmt.Println("this is indexAfter: ", indexAfter)
		fmt.Println("this is numsMap after: ", numsMap)

	}

	fmt.Println("this is maxxxx: ", max)

	fmt.Println("this is numsMap after: ", numsMap)

	return sum
}
