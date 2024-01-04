package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 1}
	fmt.Println(solve(nums))
}

func solve(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)

	for l < r {
		p := (l + r) / 2
		if nums[p] >= nums[0] {
			l = p + 1
		} else {
			if nums[p] < res {
				res = nums[p]
			}
			r = p
		}
	}

	return res
}
