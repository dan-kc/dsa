package main

import "fmt"

func main() {
	input := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(solve(input))
}
func solve(input []int) int {
	// convert to a set
	inputMap := make(map[int]bool)
	for i := 0; i < len(input); i++ {
		inputMap[input[i]] = true
	}
	longestSeqLen := 0
	for i := 0; i < len(input); i++ {
		// check if start of sequence
		if inputMap[input[i]-1] {
			// we are not at start
			continue
		}
		// we are the start of a sequence
		nextVal := input[i] + 1
		count := 1
		for inputMap[nextVal] {
			nextVal++
			count++
		}

		// check if count is significant
		if count > longestSeqLen {
			longestSeqLen = count
		}
	}
	return longestSeqLen
}
