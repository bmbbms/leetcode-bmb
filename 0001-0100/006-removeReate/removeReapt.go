package main

import "fmt"

func removeDuplicates(nums []int) int {
	var j = 0
	last_index := len(nums) - 1
	for i := 0; i < last_index; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}

	}
	nums[j] = nums[last_index]
	return j + 1
}

func main() {

	duplicates := removeDuplicates([]int{2, 3, 3, 3, 4, 4, 4, 5, 5, 6})

	fmt.Println(duplicates)
}
