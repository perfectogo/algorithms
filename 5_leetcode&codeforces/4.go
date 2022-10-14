package main

import "fmt"

func main() {
	numbers2 := []int{1, 2, 3, 4}
	numbers1 := []int{5}

	med := findMedianSortedArrays(numbers1, numbers2)
	fmt.Println(med)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	lenth := l1 + l2
	if lenth%2 == 1 {
		if nums1[l1-1] < nums2[0] {
			if l1 > l2 {
				return float64(nums1[lenth/2])
			} else {

			}

		} else {
			fmt.Println(lenth/2, ";;")
			return float64(nums2[lenth/2-l2])
		}
	}

	return 0
}
