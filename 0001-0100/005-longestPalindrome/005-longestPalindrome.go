package main

func longestPalindrome(s string) string {

	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func longestPalindrome1(s string) string {

	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i <= len(s)-1; i++ {
		for j := i; j >= 0; j-- {
			dp[i][j] = (s[i] == s[j]) && ((i-j < 3) || dp[i-1][j+1])
			if dp[i][j] && (res == "" || i-j+1 > len(res)) {
				res = s[j : i+1]
			}
		}
	}
	return res
}

func main() {
	println(longestPalindrome("abccbadef"))
	println(longestPalindrome1("abccbadef"))
}
