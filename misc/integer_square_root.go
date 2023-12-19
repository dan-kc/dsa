package main

import "fmt"

func main() {
	test := 10
	fmt.Print(Isqrt(test))
}

func ISqrt(x int) int {
	l := 0
	h := x + 1
	for l != h-1 {
		m := l + (h-l)/2
		if m*m <= x {
			l = m
		} else {
			h = m
		}
	}
	return l
}

// Or

func Isqrt(x int) int {
	l := 0
	h := x + 1
	m := -1
	for l < h {
		m = l + (h-l)/2
		if m*m == x {
			return m
		}
		if m*m > x {
			h = m - 1
		}
		if m*m < x {
			l = m + 1
		}
	}
	return m
}
