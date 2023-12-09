package main

import (
	"fmt"
)

func main() {
	var n int
	var ans int
	fmt.Scan(&n)
	for n > 0 {
		n = n / 5
		ans = ans + n
	}
	fmt.Println(ans)
}
