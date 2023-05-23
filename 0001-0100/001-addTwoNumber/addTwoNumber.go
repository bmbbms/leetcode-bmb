package main

import "fmt"

func addTwoNumber(nums []int, target int) []int {
	anothers := make(map[int]int, 0)

	for k, v := range nums {
		if idx, ok := anothers[target-v]; ok {
			return []int{idx, k}
		}
		anothers[v] = k

	}

	return nil

}

func main() {
	nums := []int{2, 5, 9}
	target := 7
	fmt.Println(addTwoNumber(nums, target))

}
