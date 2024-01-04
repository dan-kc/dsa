package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 1}
	fmt.Println(solve(nums, 1))
}

func solve(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			// we know m is in left
			if nums[m] < target || target < nums[l] {
				l = m + 1
			} else {
				r = m
			}
		} else {
			// we know m is in right
			if nums[m] < target || nums[l] <= target {
				r = m
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
