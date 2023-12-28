package main

import (
	"fmt"
)

func main() {
	test := "abbcdecdefghijk"
	fmt.Print(solve(test))
}

func solve(str string) int {
  if len(str) == 0 {
    return 0
  }
	left := 0
	right := 0
	maxlen := 1

	charidxmap := make(map[byte]int)
	for right < len(str) {
		// if right points to a char already in map
		if idx, ok := charidxmap[str[right]]; ok {
			// update l to the next val after
			left = idx + 1
		}
		// update map
		charidxmap[str[right]] = right
    // update maxlen
    maxlen = max(right - left + 1, maxlen)

		right++
	}

	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
