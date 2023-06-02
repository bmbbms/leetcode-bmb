package main

func maxArea(nums []int) int {
	maxArea := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			area := (j - i) * min(nums[i], nums[j])
			if area > maxArea {
				maxArea = area
			}

		}

	}
	return maxArea

}
func maxArea2(nums []int) int {
	length := len(nums)
	L, R := 0, length-1
	maxarea := 0
	for L < R {
		if nums[L] < nums[R] {
			left := nums[L]
			area := left * (R - L)
			if area > maxarea {
				maxarea = area
			}
			for left > nums[L+1] && L < R {
				L++
			}
			L++
		} else {
			right := nums[R]
			area := right * (R - L)
			if area > maxarea {
				maxarea = area
			}
			for right > nums[R-1] && L < R {
				R--
			}
			R--
		}

	}

	return maxarea
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	} else {
		return i2
	}

}
func main() {

	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
