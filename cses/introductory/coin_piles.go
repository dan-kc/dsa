package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		if (2*b-a)%3 == 0 && (2*a-b)%3 == 0 && (2*b-a) >= 0 && (2*a-b) >= 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
