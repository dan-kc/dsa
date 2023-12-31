package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 12
	pos := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Print(solve(target, pos, speed))
}

func solve(target int, position []int, speed []int) int {
	pair := []carInfo{}
	stack := []float32{} //stack representing the time it will reach the target
	for i := 0; i < len(position); i++ {
		pair = append(pair, carInfo{position[i], speed[i]})
	}

	// sort ascending
	sort.Slice(pair, func(i, j int) bool {
		return pair[i].pos < pair[j].pos
	})

	for i := len(pair) - 1; i >= 0; i-- {
		stack = append(stack, float32(target-pair[i].pos)/float32(pair[i].spd))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}

type carInfo struct {
	pos int
	spd int
}
