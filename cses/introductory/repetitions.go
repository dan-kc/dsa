package main

import "fmt"

// Time: O(n)
// Space: O(1)
func main() {
	var sequence string
	fmt.Scan(&sequence)
	maxRepetition := 1
	currentRepetition := 1
	for i := 1; i < len(sequence); i++ {
		if sequence[i] == sequence[i-1] {
			currentRepetition++
		} else {
			if currentRepetition > maxRepetition {
				maxRepetition = currentRepetition
			}
			currentRepetition = 1
		}
	}
	if currentRepetition > maxRepetition {
		maxRepetition = currentRepetition
	}
	fmt.Println(maxRepetition)
}

// Time: O(n)
// Space: O(n)
func main2() {
	var input string
	fmt.Scan(&input)

	repMap := make(map[rune]int)
	for _, r := range input {
		repMap[r]++
	}

	largest := 0
	for _, v := range repMap {
		if v > largest {
			largest = v
		}
	}

	fmt.Println(largest)
}
