package main

/*
*
分类

出现次数是2种的情况：

	其中一种为 出现次数为1 的字符只有1个,则 true

	另外一种为出现次数为k+1的字符只有 1个，并且另外一种次数为k

出现次数为 1种的情况

	如果字符的种类为1 则是true
	如果次数 为 1     则是true
*/
func equalFrequency(word string) bool {
	m := make(map[byte]int, 26)
	for _, char := range word {
		m[byte(char)]++
	}
	n := make(map[int]int, 26)
	for _, v := range m {
		n[v]++
	}

	if len(n) == 2 {
		keys := []int{}
		for k, v := range n {
			if k == 1 && v == 1 {
				return true
			} else {
				keys = append(keys, k)
			}
		}
		if (keys[0]-keys[1] == 1 && n[keys[0]] == 1) || (keys[1]-keys[0] == 1 && n[keys[1]] == 1) {

			return true
		}

	} else if len(n) == 1 {
		for k, v := range n {
			if k == 1 || v == 1 {
				return true
			}
		}
	}

	return false

	//if len(n) >= 3 {
	//	return false
	//}
	//if len(n) == 1 {
	//	for k, _ := range n {
	//		if k == 1 {
	//			return true
	//		}
	//
	//	}
	//
	//}
	//keys := []int{}
	//if len(n) == 2 {
	//	for k := range n {
	//		keys = append(keys, k)
	//
	//	}
	//	if keys[0]-keys[1] == 1 || keys[1]-keys[0] == 1 {
	//		return true
	//
	//	} else {
	//		return false
	//	}
	//
	//}
	return false
}
func main() {

}
