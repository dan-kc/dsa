package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fastPow2(n))
}

var MOD = 1e9 + 7

func fastPow2(x int) int {
	res := 1
	base := 2
	for x != 0 {
		if isOdd(x) {
			res = (res * base) % int(MOD)
		}
		base = (base * base) % int(MOD)
		x = int(math.Floor(float64(x) / 2))
	}
	return res
}

func isOdd(x int) bool {
	return x%2 == 1
}

// OR

func main() {
	var n int
	fmt.Scan(&n)
	a := fastPow(2, n)
	fmt.Println(a)
}

func fastPow(a, n int) int {
	var MOD = 1e9 + 7
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	r := fastPow(a, n/2)
	if n%2 == 0 {
		return r * r % int(MOD)
	}
	return r * r * a % int(MOD)
}
