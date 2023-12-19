package main

import "fmt"

func main() {
	test := []int{1, 3, 6, 4, 2}
	sort(test)
	fmt.Println(test)
}

func sort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
