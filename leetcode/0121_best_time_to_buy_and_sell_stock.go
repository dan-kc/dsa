package main

import (
	"fmt"
	"math"
)

func main() {
	test := []int{7, 1, 5, 3, 6, 4, 9, 0, 12}
	fmt.Print(solve(test))
}

func solve(arr []int) int {
	l := 0
	r := 1
	max := 0
	for r < len(arr) {
		if arr[r] > arr[l] {
			profit := arr[r] - arr[l]
			max = int(math.Max(float64(profit), float64(max)))
		} else {
			l = r
		}
		r++
	}
	return max
}
