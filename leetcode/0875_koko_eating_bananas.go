package main

import "fmt"

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(solve(piles, h))
}

func solve(piles []int, h int) int {
	max := max(piles)
	l := 1
	r := max + 1
	var k int
	for l < r {
		k = (l + r) / 2
		if isPossible(piles, h, k) {
			r = k
		} else {
			l = k + 1
		}
	}
	return k
}

func isPossible(piles []int, h, k int) bool {
	var count int
	for i := 0; i < len(piles); i++ {
		count += (piles[i] / k) + 1
	}
	return count <= h
}

func max(input []int) int {
	max := -1
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}
