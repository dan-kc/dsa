package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := n * (n + 1) / 2

	if sum%2 > 0 {
		fmt.Println("NO SOLUTION")
		return
	}

	var left, right []int
	if n%4 == 0 {
		for i := 1; i < n+1; i += 4 {
			left = append(left, i, i+3)
			right = append(right, i+1, i+2)
		}
	} else {
		left = append(left, 1, 2)
		right = append(right, 3)
		for i := 4; i < n+1; i += 4 {
			left = append(left, i, i+3)
			right = append(right, i+1, i+2)
		}
	}

	fmt.Println("YES")

	fmt.Println(len(left))
	for i := 0; i < len(left); i++ {
		fmt.Printf("%d ", left[i])
	}

	fmt.Println()
	fmt.Println(len(right))
	for i := 0; i < len(right); i++ {
		fmt.Printf("%d ", right[i])
	}
}
