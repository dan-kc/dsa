package main

import (
	"fmt"
	"sort"
)

func main() {
	var input string
	fmt.Scan(&input)
	runes := []rune(input)
	// sort input
	sort.Slice(runes, func(i, j int) bool {
		return runes[j] < runes[i]
	})

	perms := make(map[string]bool)

	perms[string(runes)] = true
	for nextPermutation(runes) {
		perms[string(runes)] = true
	}

	fmt.Println(len(perms))
	for k, _ := range perms {
		fmt.Println(k)
	}
}

func nextPermutation(arr []rune) bool {
	if len(arr) == 0 || arr == nil {
		return false
	}

	// find breakpoint
	breakpoint := -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] { // if not perfectly sorted
			breakpoint = i
			break
		}
	}
	if breakpoint == -1 {
		return false
	}

	// can change to binary
	// find sm element index in suffix
	sm := len(arr) - 1
	for i := len(arr) - 2; i > breakpoint; i-- {
		if arr[i] < arr[sm] {
			sm = i
		}
	}

	// swap elements at sm and breakpoint
	arr[breakpoint], arr[sm] = arr[sm], arr[breakpoint]

	// sort the suffix array.
	// We already know that it is ordered.
	// So we just have to reverse it
	left := breakpoint + 1
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return true
}
