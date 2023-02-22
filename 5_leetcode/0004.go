package main

import (
	"fmt"
	"math"
)

var (
	nums1 = []int{1, 3}
	nums2 = []int{2}
	nums3 = []int{1, 2}
	nums4 = []int{3, 4}
)

func main() {
	fmt.Println(findMedianSortedArrays(nums3, nums4))
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	var m, n int = len(nums1), len(nums2)

	if m > n {
		nums1, nums2 = nums2, nums1
		n, m = m, n
	}

	var (
		imin     = 0
		imax     = m
		half_len = (m + n + 1) / 2
	)

	for imin <= imax {
		i := (imin + imax) / 2
		j := half_len - i

		if i < imax && nums1[i] < nums2[j-1] {
			imin = i + 1
		} else if i > imin && nums2[j] < nums1[i-1] {
			imax = i - 1
		} else {
			var max_left float64

			if i == 0 {
				max_left = float64(nums2[j-1])
			} else if j == 0 {
				max_left = float64(nums1[i-1])
			} else {
				max_left = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return max_left
			}

			var min_right float64

			if i == m {
				min_right = float64(nums2[j])
			} else if j == n {
				min_right = float64(nums1[i])
			} else {
				min_right = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (max_left + min_right) / 2.0
		}
	}
	return 0
}
