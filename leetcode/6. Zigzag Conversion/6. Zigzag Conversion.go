package main

import "fmt"

func convert(word string, rows int) string {
	var index, diagonalIndex int
	result := make([]byte, 0, len(word))

	if rows == 1 {
		return word
	}

	for i := range rows {
		index = i
		for index < len(word) {
			result = append(result, word[index])
			diagonalIndex = index + 2*(rows-1-i)

			if i > 0 && i < rows-1 && diagonalIndex < len(word) {
				result = append(result, word[diagonalIndex])
			}
			index += max(1, 2*rows-2)
		}
	}
	return string(result)
}

func main() {
	word := "PAYPALISHIRING"
	rows := 3
	result := convert(word, rows)
	fmt.Println(result)
}

// P   A   H   N
// A P L S I I G
// Y   I   R
