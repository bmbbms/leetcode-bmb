package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	result := []int{}
	j := 0
	for i := 0; i < len1; i++ {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	for j < len2 {
		result = append(result, nums2[j])
		j++
	}
	fmt.Println(result)
	return 30.0

}
func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})

}
