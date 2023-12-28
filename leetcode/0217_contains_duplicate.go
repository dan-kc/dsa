package leetcode

import (
	"slices"
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
