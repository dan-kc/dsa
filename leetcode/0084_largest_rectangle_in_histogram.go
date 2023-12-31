package main

import (
	"fmt"
)

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(solve(heights))
}

func solve(heights []int) int {
	// find left most limit for all bars
	lLimits := make([]int, len(heights))
	lStack := make([]int, 0, len(heights))
	for i := 0; i < len(heights); i++ {
		lLimit := 0
		// compare the top(s) of the stack to the current value
		for len(lStack) != 0 {
			if heights[i] > heights[lStack[len(lStack)-1]] {
				lLimit = lStack[len(lStack)-1] + 1
				break
			}
			// pop
			lStack = lStack[:len(lStack)-1]
		}
		lLimits[i] = lLimit
		lStack = append(lStack, i)
	}

	// find the right limits
	rLimits := make([]int, len(heights))
	rStack := make([]int, 0, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		rLimit := len(heights) - 1
		// compare the top(s) of the stack to the current value
		for len(rStack) != 0 {
			if heights[i] > heights[rStack[len(rStack)-1]] {
				rLimit = rStack[len(rStack)-1] - 1
				break
			}
			// pop
			rStack = rStack[:len(rStack)-1]
		}
		rLimits[i] = rLimit
		rStack = append(rStack, i)
	}

	// calc max areas:
	areas := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		maxWidth := (i - lLimits[i]) + (rLimits[i] - i) + 1
		areas[i] = heights[i] * maxWidth
	}
	return Max(areas)
}

func Max(values []int) int {
	max := 0
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}
