package main

func main() {

}

func removeDuplicates(nums []int) int {

	var lenth int = len(nums)

	for i := 0; i < lenth-1; i++ {
		for j := i + 1; j < lenth-1; j++ {
			if nums[i] == nums[j] {

			}
		}
	}
	return 0
}

func delete_at_index(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
