package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	if length < 4 {
		return [][]int{}

	}
	result := [][]int{}
	for i := 0; i < length-4; i++ {

		for j := i + 1; j < length-3; j++ {
			L, R := j+1, length-1
			for L < R {
				if nums[i]+nums[j]+nums[L]+nums[R] > target {
					R--
				} else if nums[i]+nums[j]+nums[L]+nums[R] < target {
					L++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[L], nums[R]})
					for nums[L] == nums[L+1] && L < R {
						L++
					}
					for nums[R] == nums[R-1] && L < R {
						R--
					}
					L++
					R--
				}
			}

			for nums[j] == nums[j+1] && j < length-3 {
				j++
			}
		}
		for nums[i] == nums[i+1] && i < length-3 {
			i++
		}

	}
	return result

}
func main() {
	fmt.Printf("%#v", fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
