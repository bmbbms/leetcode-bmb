package main

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	romanmaps := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	tmp := 0
	for i := 0; i < len(s)-1; i++ {
		if romanmaps[s[i]] > romanmaps[s[i+1]] {
			result = result + tmp + romanmaps[s[i]]
			tmp = 0
		} else if romanmaps[s[i]] == romanmaps[s[i+1]] {
			tmp = tmp + romanmaps[s[i]]
		} else {
			tmp = tmp + romanmaps[s[i]]
			result = result - tmp
			tmp = 0
		}

	}

	return result + tmp + romanmaps[s[len(s)-1]]

}

func intToRoman(nums int) string {
	strs := []string{"I", "V", "X", "L", "C", "D", "M"}
	mapsromans := make(map[int]map[int]string, 4)
	for i := 0; i < 3; i++ {
		mapsromans[i] = make(map[int]string, 10)
		mapsromans[i][0] = ""
		mapsromans[i][1] = strs[2*i]
		mapsromans[i][2] = strs[2*i] + mapsromans[i][1]
		mapsromans[i][3] = strs[2*i] + mapsromans[i][2]
		mapsromans[i][4] = strs[2*i] + strs[2*i+1]
		mapsromans[i][5] = strs[2*i+1]
		mapsromans[i][6] = strs[2*i+1] + strs[2*i]
		mapsromans[i][7] = mapsromans[i][6] + strs[2*i]
		mapsromans[i][8] = mapsromans[i][7] + strs[2*i]
		mapsromans[i][9] = strs[2*i] + strs[2*i+2]

	}
	mapsromans[3] = map[int]string{0: "", 1: "M", 2: "MM", 3: "MMM"}

	//map[int]map[int]string{
	//	0: {
	//		1: "I",
	//		2: "II",
	//		3: "III",
	//		4: "IV",
	//		5: "V",
	//		6: "VI",
	//		7: "VII",
	//		8: "VIII",
	//		9: "IX",
	//	},
	//
	//}

	return mapsromans[3][nums/1000] + mapsromans[2][nums/100%10] +
		mapsromans[1][nums/10%10] + mapsromans[0][nums%10]

}

func main() {
	println(romanToInt("CMI"))

	println(intToRoman(909))
}
