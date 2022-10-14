package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if (i != j) && (num1+num2 == target) {
				return []int{i, j}
			}
		}
	}
	return nil
}
