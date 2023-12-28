package main

import (
	"fmt"
)

func main() {
	str1 := "oabo"
	str2 := "eidbaooo"
	fmt.Print(solve(str1, str2))
}

func solve(str1, str2 string) bool {
	// make map of str1
	str1Count := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		str1Count[str1[i]]++
	}

	// make map of starting window
	windowCount := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		windowCount[str2[i]]++
	}
	if isEqual(str1Count, windowCount) {
		return true
	}

	// define left and right of our window
	l := 1
	r := len(str1)
	for r < len(str2) {
		windowCount[str2[r]]++
		windowCount[str2[l-1]]--
		if isEqual(str1Count, windowCount) {
			return true
		}
		r++
		l++
	}
	return false
}

func isEqual(count1, count2 map[byte]int) bool {
	for k, v := range count2 {
		if count1[k] != v {
			return false
		}
	}
	return true
}
