package leetcode

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
		// This is leaky because of utf-8. e.g. s="abcâ" t="abcde"
		// It still gets caught later though
	}
	sCount := make(map[rune]int)
	tCount := make(map[rune]int)
	for _, r := range s {
		sCount[r]++
	}
	for _, r := range t {
		tCount[r]++
	}
	for _, r := range s {
		if tCount[r] != sCount[r] {
			return false
		}
	}
	return true
}

func ValidAnagramAlt(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var freq [26]int
	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}
	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}
	return true
}
