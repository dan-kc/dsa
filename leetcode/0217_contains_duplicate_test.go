package leetcode_test

import (
	"reflect"
	"testing"

	"github.com/dan-kc/dsa/leetcode"
)

func TestContainsDuplicateTest(t *testing.T) {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, v := range testCases {
		result := leetcode.ContainsDuplicate(v.input)
		if !reflect.DeepEqual(result, v.expected) {
			t.Errorf("For inputs %d,  expected %v, but got %v", v.input, v.expected, result)
		}
	}
}
