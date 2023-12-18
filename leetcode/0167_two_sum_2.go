package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	target := 7
	fmt.Print(TwoSum(input, target))
}

func TwoSum(input []int, target int) (int, int) {
	left := 0
	right := len(input) - 1
	for right != left+1 {
		if input[left]+input[right] > target {
			right--
			continue
		}
		if input[left]+input[right] < target {
			left++
			continue
		}
		break
	}
	return left + 1, right + 1
}
