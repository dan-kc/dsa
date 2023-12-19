package main

import "fmt"

func main() {
	test := []int{2, 8, 13, 22, 38, 45, 53, 87, 99}
	fmt.Print(binarySearch(test, 45))
}

// a  binary search only operates on an ordered array
func binarySearch(input []int, target int) int {
	// l inclusive, h exclusive (consider target at end)
	l := 0
	h := len(input)
	for l < h {
		m := l + (h-l)/2
		if target == input[m] {
			return m
		}
		if target < input[m] {
			h = m - 1
		}
		if target > m {
			l = m + 1
		}
	}
	return -1
}
