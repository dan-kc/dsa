package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	ans := solve(n)
	for i := 0; i < len(ans); i++ {
		fmt.Println()
		for j := 0; j < n; j++ {
			fmt.Printf("%b", ans[i][j])
		}
	}
}

func solve(n int) [][]byte {
	if n == 1 {
		return [][]byte{{0}, {1}}
	}
	prev := solve(n - 1)

	next := [][]byte{}
	len := len(prev)
	for i := 0; i < len; i++ {
		entry := append([]byte{0}, prev[i]...)
		next = append(next, entry)
	}
	for i := len - 1; i >= 0; i-- {
		entry := append([]byte{1}, prev[i]...)
		next = append(next, entry)
	}
	return next
}
