package main

import "fmt"

func main() {
	test := []int{2, 3, 5, 2, 2, 10, 9, 5}
	sort(test, 0, len(test)-1)
	fmt.Print(test)
}

func sort(arr []int, low, high int) {
	if low < high { // there are at least two elements
		p := partition(arr, low, high)
		sort(arr, low, p-1)
		sort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int { // returns the new index of the pivot
	p := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= p {
			i++ // index of the first element that is greater than p
		}
		for arr[j] > p {
			j-- // index of the first element that is less than or equal to p
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}
