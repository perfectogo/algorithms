package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1, 3, 9, 11}, 4))
}

func searchInsert(nums []int, target int) int {

	lenth := len(nums)

	for i := 0; i < lenth; i++ {

		if nums[0] > target {
			return 0
		}

		if nums[lenth-1] < target {
			return lenth
		}

		if nums[i] == target {
			return i

		}

		if i+1 < lenth {
			if nums[i] < target && nums[i+1] > target {
				return i + 1
			}
		}
	}

	return 0
}

/*
func searchInsert(nums []int, target int) int {

    for i, n := range nums {

		if n == target || n > target {

			return i

		}

	}

	return len(nums)

}

func searchInsert(nums []int, target int) int {
    start, end := 0, len(nums)-1
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return end + 1
}
*/
