package main

import (
	"fmt"
)

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	period := 2*numRows - 2
	result := []byte{}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= (len(s)+i-1)/period; j++ {
			if i == 0 || i == numRows-1 {
				if period*j+i < len(s) {
					result = append(result, s[period*j+i])
				}

				continue
			} else {
				if period*j+i < len(s) {
					result = append(result, s[period*j+i])
				}

				if period*(j+1)-i < len(s) {
					result = append(result, s[period*(j+1)-i])
				}

			}

		}
	}
	return string(result)

}

func convert1(s string, numrows int) string {

	result, down, up := make([][]byte, numrows), 0, numrows-2

	for i := 0; i < len(s); {
		if down != numrows {
			result[down] = append(result[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			result[up] = append(result[up], byte(s[i]))
			up--
			i++
		} else {
			down = 0
			up = numrows - 2
		}

	}

	// solution := make([]byte, len(s)) must add length when use append，if not，the result will be not correct
	solution := make([]byte, 0, len(s))

	for _, row := range result {
		for _, item := range row {
			solution = append(solution, item)
		}

	}
	d := string(solution)
	fmt.Println(len(d))
	fmt.Println(string(d))
	return string(solution)

}

func convert3(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}

func main() {

	conv1 := convert("PAYPALISHIRING", 3)
	conv2 := convert1("PAYPALISHIRING", 3)
	conv3 := convert3("PAYPALISHIRING", 3)
	target := "PAHNAPLSIIGYIR"
	fmt.Println(conv1 == target)
	fmt.Println(conv2 == target)
	fmt.Println(conv3 == target)
	fmt.Println(target)
	fmt.Println("=====")

	for _, s := range conv1 {
		fmt.Printf("%c", s)
	}
	fmt.Println()
	for _, s := range conv2 {
		fmt.Printf("%c", s)
	}

}
