package main

import "fmt"

func move(from, to, depth int) {
	if depth == 1 {
		fmt.Printf("%d %d\n", from, to)
		return
	}

	other := 6 - from - to
	move(from, other, depth-1)
	fmt.Printf("%d %d\n", from, to) // move bottom from to
	move(other, to, depth-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d\n", (1<<n)-1) // (1<<n) is 2^n lmao
	move(1, 3, n)
}
