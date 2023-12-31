package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(solve(3))
}

func solve(n int) []string {
	var stack []string
	var res []string
	backtrack(n, 0, 0, &stack, &res)
	return res
}

func backtrack(n int, openN int, closedN int, stack *[]string, res *[]string) {
	// base case / terminating condition
	if openN == n && closedN == n {
		*res = append(*res, strings.Join(*stack, ""))
	}
	if openN < n {
		*stack = append(*stack, "(")
		backtrack(n, openN+1, closedN, stack, res)
		*stack = (*stack)[:len(*stack)-1]
	}
	if closedN < openN {
		*stack = append(*stack, ")")
		backtrack(n, openN, closedN+1, stack, res)
		*stack = (*stack)[:len(*stack)-1]
	}
}
