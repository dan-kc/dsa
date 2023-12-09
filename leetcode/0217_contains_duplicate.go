package leetcode

import (
	"reflect"
	"slices"
	"testing"
)

// Time: O(nlog(n))
// Space: O(1)
func ContainsDuplicate(nums []int) bool {
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// Time: O(n^2)
// Space: O(1)
func ContainsDuplicateAlt(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Time: O(n)
// Space: O(n)
func ContainsDuplicateAlt2(nums []int) bool {
	seen := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if seen[i] {
			return true
		}
		seen[i] = true
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
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
