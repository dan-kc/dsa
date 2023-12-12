package main

import "fmt"

// Time: O(n)
func main() {
	var n int
	fmt.Scan(&n)

	ans := []int{}
	for i := 0; i < n; i++ {
		var y, x int
		fmt.Scan(&y)
		fmt.Scan(&x)

		var val int
		max := max(y, x)
		if max%2 == 0 {
			if y < x {
				// 2,4
				// 11
				//
				val = (max-1)*(max-1) + y
			} else {
				// 4, 2
				// 15
				val = (max-1)*(max-1) + y + (max - x)
			}
		} else {
			if y < x {
				// 2,5
				// 24
				val = (max-1)*(max-1) + x + (max - y)
			} else {
				// 5, 2
				// 18
				val = (max-1)*(max-1) + x
			}
		}
		ans = append(ans, val)
	}

	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}
