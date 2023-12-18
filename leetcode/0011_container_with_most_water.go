package main

import (
	"fmt"
)

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Print(solve(input))
}

func solve(nums []int) int {
	left := 0
	right := len(nums) - 1
	largest := 0
	for left < right {
		height := min(nums[left], nums[right])
		width := right - left
		area := height * width
		if area > largest {
			largest = area
		}

		// we know that the next area we calculate will have a smaller width.
		// therefore we know it needs a bigger height

		if nums[left] <= nums[right] {
			left++
			continue
		}
		right--
	}
	return largest
}
