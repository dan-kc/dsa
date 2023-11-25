package leetcode_test

import (
	"reflect"
	"testing"

	"github.com/dan-kc/dsa/leetcode"
)

func TestValidAnagram(t *testing.T) {
	testCases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		//5 abcâ
		//5 abcde
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"abcâ", "abcde", false},
	}
	for _, v := range testCases {
		result := leetcode.ValidAnagram(v.input1, v.input2)
		if !reflect.DeepEqual(result, v.expected) {
			t.Errorf("For inputs %s and %s,  expected %v, but got %v", v.input1, v.input2, v.expected, result)
		}
	}
}
