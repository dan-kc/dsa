// Time: O(n)
// Space O(n)
package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if val, found := m[target-num]; found {
			return []int{val, idx}
		}
		m[num] = idx
	}
	return nil
}

// Time: O(n^2)
// Space: O(1)
// func TwoSum(nums []int, target int) (int, int) {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return i, j
// 			}
// 		}
// 	}
// 	return 0, 0
// }
