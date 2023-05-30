package main

import "math"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	for x%10 == 0 {
		x = x / 10
	}
	var result int
	for x != 0 {
		num := x % 10
		x = x / 10
		result = result*10 + num
	}

	if result > int(math.Pow(2, 31))-1 || result < int(math.Pow(-2, 31)) {
		return 0
	}
	return result
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	palindrome := reverse(x)
	if palindrome == x {
		return true
	}
	return false
}
