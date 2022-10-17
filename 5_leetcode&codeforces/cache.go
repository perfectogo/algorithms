package main

// //////////////// 1 for time
func zeroArray(array []int) float64 {
	if len(array)%2 == 0 {
		return (float64(array[len(array)/2]) + float64(array[(len(array)-1)/2])) / 2
	} else {
		return float64(array[len(array)/2])
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	i1, i2 := 0, 0

	var cur_value int
	var prev_value int

	if len(nums2) == 0 {
		zeroArray(nums1)
	}

	if len(nums1) == 0 {
		zeroArray(nums2)
	}

	for {

		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] <= nums2[i2] {
				prev_value = cur_value
				cur_value = nums1[i1]
				i1 += 1
			} else {
				prev_value = cur_value
				cur_value = nums2[i2]
				i2 += 1
			}
		} else {

			if i1 == len(nums1) {
				prev_value = cur_value
				cur_value = nums2[i2]
				i2 += 1
			} else if i2 == len(nums2) {
				prev_value = cur_value
				cur_value = nums1[i1]
				i1 += 1
			}
		}

		if i1+i2-1 == (len(nums1)+len(nums2))/2 {

			if (len(nums1)+len(nums2))%2 != 0 {
				return float64(cur_value)
			} else {
				return (float64(cur_value) + float64(prev_value)) / 2
			}
		}
	}
}

/////////////////// 2 for memory

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	index1 := (m + n - 1) / 2
	needNext := (m+n)%2 == 0

	var current, j, k int
	for j, k = -1, -1; j+k+2 <= index1; {
		current = getNextItem(nums1, &j, nums2, &k)
	}

	if needNext {
		next := getNextItem(nums1, &j, nums2, &k)

		return (float64(current) + float64(next)) / 2
	}

	return float64(current)
}

func getNextItem(nums1 []int, index1 *int, nums2 []int, index2 *int) int {
	m, n := len(nums1), len(nums2)

	if *index1 >= m-1 {
		*index2++
		return nums2[*index2]
	}

	if *index2 >= n-1 {
		*index1++
		return nums1[*index1]
	}

	if nums1[*index1+1] < nums2[*index2+1] {
		*index1++
		return nums1[*index1]
	}

	*index2++
	return nums2[*index2]
}
