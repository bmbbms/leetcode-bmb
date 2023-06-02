package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	result := make([][]int, 0, 3)

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] > 0 {
				break
			}
			for k := j + 1; k < len(nums); k++ {

				if nums[i]+nums[j]+nums[k] > 0 {
					break
				} else if nums[i]+nums[j]+nums[k] == 0 {
					fmt.Println(nums[i], nums[j], nums[k])
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}

		}

	}
	var j = 0
	for i := 0; i < len(result)-1; i++ {
		if result[i][0] == result[i+1][0] && result[i][1] == result[i+1][1] && result[i][2] == result[i+1][2] {
			continue
		} else {
			result[j] = result[i]
			j++
		}
	}
	if len(result) == 0 {
		return [][]int{}
	}
	return result[:j+1]

}

func threeSums(nums []int) [][]int {
	length := len(nums)
	result := make([][]int, 0, 3)
	if length < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {

		L, R := i+1, length-1
		for L < R {
			if nums[i]+nums[L]+nums[R] < 0 {
				L = L + 1
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R = R - 1
			} else {
				result = append(result, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L = L + 1
				}
				for L < R && nums[R] == nums[R-1] {
					R = R - 1
				}
				L = L + 1
				R = R - 1
			}
		}

		for nums[i] == nums[i+1] && i < length-2 {
			i++
		}

	}

	if len(result) == 0 {
		return [][]int{}
	}
	return result
}

func main() {
	sum := threeSums([]int{0, 0, 0})
	fmt.Println(sum)
}
