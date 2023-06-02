package main

var maps []string
var results []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	maps = []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	results = []string{}
	getAllresult(digits, 0, "")
	return results

}

func getAllresult(digits string, index int, resultOne string) {
	if index == len(digits) {
		results = append(results, resultOne)
	} else {
		letters := maps[digits[index]-'0']

		for i := 0; i < len(letters); i++ {
			getAllresult(digits, index+1, resultOne+string(letters[i]))
		}

	}

}

func main() {
	letterCombinations("23")
}
