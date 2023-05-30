package main

import (
	"errors"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	flag := ""

	if s[0] == '-' {
		flag = "-"
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	lastIndex := -1
	for char := range s {
		if s[char] <= '9' && s[char] >= '0' {
			lastIndex = char
			continue
		} else {
			break
		}
	}

	if lastIndex > -1 {
		s = flag + s[:lastIndex+1]
	} else {
		return 0
	}
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			if flag == "" {
				return 2<<30 - 1
			} else {
				return -(2 << 30)
			}
		}
		return 0
	}

	//if i > (2<<30)-1 {
	//	return 2<<30 - 1
	//}
	//
	//if i < -(2 << 30) {
	//	return -(2 << 30)
	//}

	return int(i)
}
func main() {
	println(myAtoi("20000000000000000000"))
}
