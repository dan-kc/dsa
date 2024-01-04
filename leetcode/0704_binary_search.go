package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := -1
	fmt.Print(solve(nums, target))
}

func solve(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			r = m
			continue
		}
		l = m + 1
	}

	return -1
}
