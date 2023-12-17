package main

import (
	"bytes"
	"fmt"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	fmt.Print(solve(input))
}

func solve(input string) bool {
	left := 0
	right := len(input) - 1
	for left < right {
		for !isValid(input[right]) {
			right--
		}
		for !isValid(input[left]) {
			left++
		}
		if bytes.ToLower([]byte{input[left]})[0] != bytes.ToLower([]byte{input[right]})[0] {
			return false
		}
		right--
		left++
	}
	return true
}

func isValid(b byte) bool {
	switch b {
	case byte(' '), byte(','), byte(';'), byte(':'):
		return false
	}
	return true
}
