package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		if !isValidGroup(board[i]) {
			return false
		}
	}

	// Check columns
	for i := 0; i < 9; i++ {
		column := make([]byte, 9)
		for j := 0; j < 9; j++ {
			column[j] = board[j][i]
		}
		if !isValidGroup(column) {
			return false
		}
	}

	// Check 3x3 sub-boxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			// we're at the start of a 3x3 box
			subbox := make([]byte, 9)
			index := 0
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					subbox[index] = board[x][y]
					index++
				}
			}
			if !isValidGroup(subbox) {
				return false
			}
		}
	}

	return true
}

func isValidGroup(unit []byte) bool {
	seen := make(map[byte]bool)
	for _, digit := range unit {
		if digit != '.' {
			if seen[digit] {
				return false
			}
			seen[digit] = true
		}
	}
	return true
}

func main() {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Example 1:", isValidSudoku(board1)) // Output: true
	fmt.Println("Example 2:", isValidSudoku(board2)) // Output: false
}
