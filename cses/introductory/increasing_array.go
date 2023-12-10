package main

import "fmt"

// Time O(n)
// Space O(n)
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var moves int
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			moves = moves + arr[i-1] - arr[i]
			arr[i] = arr[i-1]
		}
	}
	fmt.Println(moves)
}
