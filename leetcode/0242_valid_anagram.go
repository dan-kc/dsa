package leetcode

import (
	"reflect"
	"testing"
)

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
		// We increment freq for characters of s and decrement t
	}
	for idx := 0; idx < len(freq); idx++ {
		// We now check if all of the map values equal zero, as they should if we have an anagram.
		if freq[idx] != 0 {
			return false
		}
	}
	return true
}

func TestValidAnagram(t *testing.T) {
	testCases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"abcâ", "abcde", false},
	}
	for _, v := range testCases {
		result := leetcode.ValidAnagramAlt(v.input1, v.input2)
		if !reflect.DeepEqual(result, v.expected) {
			t.Errorf("For inputs %s and %s,  expected %v, but got %v", v.input1, v.input2, v.expected, result)
		}
	}
}
