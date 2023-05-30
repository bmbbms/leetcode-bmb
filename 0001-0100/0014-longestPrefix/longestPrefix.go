package main

func longestCommonPrefix(strs []string) string {
	nums := len(strs)

	if nums == 0 {
		return ""
	}

	if nums == 1 {
		return strs[0]
	}

	base := strs[0]
	var i int
OuterLoop:
	for i = 0; i < len(base); i++ {
		for j := 1; j < nums; j++ {
			if i > len(strs[j]) {
				break OuterLoop
			}
			if base[i] != strs[j][i] {
				break OuterLoop
			}
		}

	}

	return base[:i]

}
func main() {

	strs := []string{"flow", "fltes", "fl"}

	println(longestCommonPrefix(strs))

}
