package main

import (
	"fmt"
)

func main() {
	test := "ABAB"
	k := 2
	fmt.Print(solve(test, k))
}

func solve(str string, k int) int {
	left := 0
	right := 0
	maxLength := 0
	count := make(map[byte]int)

	for right < len(str) {
		count[str[right]]++
		currentLength := right - left + 1
		largest := largestValue(count)

		// Check if the number of replacements needed exceeds k
		if currentLength-largest > k {
			count[str[left]]--
			left++
		}

		// Update the maximum length
		maxLength = max(maxLength, right-left+1)

		right++
	}

	return maxLength
}

func largestValue(count map[byte]int) int {
	max := 0
	for _, v := range count {
		if v > max {
			max = v
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
