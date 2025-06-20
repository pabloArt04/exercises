package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var max, lenSub, start int
	index := make(map[byte]int)

	for i := range len(s) {
		prevIndex, ok := index[s[i]]

		if ok && prevIndex >= start {
			lenSub = i - prevIndex
			start = prevIndex + 1
		} else {
			lenSub++
		}

		index[s[i]] = i

		if lenSub > max {
			max = lenSub
		}
	}

	return max
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}
