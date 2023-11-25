package leetcode

// Time: O(n^2)
// Space: O(1)
func TwoSumAlt(nums []int, target int) (int, int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return 0, 0
}

func TwoSum(nums []int, target int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums)-1; i++ {
		s := target - nums[i]
		if m[s] != 0 {
			return i, m[s]
		}
	}
	return 0, 0
}
