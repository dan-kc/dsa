package main

import "fmt"

func main() {
	test := []int{3, 5, 6, 8, 2, 1}
	fmt.Print(Sort(test))
}

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	m := len(arr) / 2
	a1 := Sort(arr[:m])
	a2 := Sort(arr[m:])

	return merge(a1, a2)
}

func merge(a, b []int) (r []int) {
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[i])
			j++
		}
	}
	// we've added all of the elements for one of the arrays

	for i < len(a) {
		r = append(r, a[i])
		i++
	}

	for j < len(b) {
		r = append(r, b[j])
		j++
	}

	return
}
