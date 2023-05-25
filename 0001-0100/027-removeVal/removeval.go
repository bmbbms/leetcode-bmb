package main

import "fmt"

func removeElement(nums []int, val int) int {

	j := 0

	for i, _ := range nums {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}

	}
	return j

}

func main() {
	nums := []int{3, 2, 2, 83, 43, 3}
	element := removeElement(nums, 3)

	for i := 0; i < element; i++ {
		fmt.Println(nums[i])
	}

	a := [3]int{0, 1, 2} //数组

	for i, _ := range a { // index、value 都是从a数组的复制品中取出。

		if i == 0 { // 在修改前先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}

	}

	fmt.Println(a) // 输出 [100, 101, 102]。

}
