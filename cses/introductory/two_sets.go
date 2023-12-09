package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := n * (n + 1) / 2
	if sum%2 != 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	var arr1 []int
	var arr2 []int
	if n%2 == 0 {
		arr1, arr2 = algo(n, 1)
	} else {
		arr1, arr2 = algo(n, 4)
	}
	fmt.Println(len(arr1))
	fmt.Println(arr1)
	fmt.Println(len(arr2))
	fmt.Println(arr2)
}

func algo(end, start int) ([]int, []int) {
	var arr1 []int
	var arr2 []int
	arr1 = append(arr1, 1, 2)
	arr2 = append(arr2, 3)
	left := true
	for start < end {
		if left {
			arr1 = append(arr1, start, end)
		} else {
			arr2 = append(arr2, start, end)
		}
		start++
		end--
		left = !left
	}
	return arr1, arr2
}
