package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(0)
	for k := 2; k < n+1; k++ {
		// we kinda divide both by 2 both due to double counting
		total := k * k * ((k * k) - 1) / 2
		attacks := 4 * 2 * (k - 1) * (k - 2) / 2
		fmt.Println(total - attacks)
	}
}
