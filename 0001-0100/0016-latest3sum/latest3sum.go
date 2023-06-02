package main

import "sort"

func three3sum(nums []int, target int) int {
	length := len(nums)
	result := 10000000
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		L, R := i+1, length-1
		for L < R {
			expected := nums[i] + nums[L] + nums[R]
			if expected < target {
				L = L + 1
				if abs(target-expected) < abs(result) {
					result = expected
				}

			} else if expected > target {
				R = R - 1
				if abs(target-expected) < abs(result) {
					result = expected
				}
			} else {
				return target
			}
		}
		for nums[i] == nums[i+1] && i < length-2 {
			i++
		}
	}
	return result
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num

}
func main() {

}
