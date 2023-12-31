package main

import (
	"fmt"
)

func main() {
	fmt.Print(solve([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

type Entry struct {
	Value int
	Index int
}

func solve(arr []int) []int {
	res := make([]int, len(arr))
	stack := make([]Entry, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		entry := Entry{Value: arr[i], Index: i}
		for j := len(stack) - 1; j >= 0; j-- {
			count := 0
			if stack[j].Value < entry.Value {
				days := entry.Index - stack[j].Index
				res[stack[j].Index] = days
				count++
			}
			stack = stack[:len(stack)-count]
		}
		stack = append(stack, entry)
	}
	return res
}
