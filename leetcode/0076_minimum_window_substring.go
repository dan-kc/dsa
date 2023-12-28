package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Print(solve(s, t))
}

// Time: O(n)
// Space: O(1)
func solve(s, t string) string {
	// create count map of target string
	targetCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	windowCount := make(map[byte]int)
	ans := ""
	matches := 0
	left, right := 0, 0

	for right < len(s) {
		// add right pos to windowCount map
		// update match count
		windowCount[s[right]]++
		if windowCount[s[right]] == targetCount[s[right]] {
			matches++
		}

		for matches == len(t) {
			// update ans
			currWindow := s[left : right+1]
			if len(currWindow) < len(ans) || ans == "" {
				ans = currWindow
			}
			// check current in map
			// update match
			if windowCount[s[left]] == targetCount[s[left]] {
				matches--
			}
			windowCount[s[left]]--

			left++
		}

		right++
	}

	return ans
}
