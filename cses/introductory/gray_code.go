package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	ans := solve(n)
	for i := 0; i < len(ans); i++ {
		fmt.Println()
		for j := 0; j < n; j++ {
			if ans[i][j] {
				fmt.Printf("%d", 1)
				continue
			}
			fmt.Printf("%d", 0)
		}
	}
}

func solve(n int) [][]bool {
	if n == 1 {
		return [][]bool{{false}, {true}}
	}
	prev := solve(n - 1)

	next := [][]bool{}
	len := len(prev)
	for i := 0; i < len; i++ {
		entry := append([]bool{false}, prev[i]...)
		next = append(next, entry)
	}
	for i := len - 1; i >= 0; i-- {
		entry := append([]bool{true}, prev[i]...)
		next = append(next, entry)
	}
	return next
}
