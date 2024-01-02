package main

import (
	"fmt"
)

func main() {
	input := [9][9]byte{
		[9]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[9]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[9]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[9]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[9]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[9]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[9]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[9]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[9]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(solve(input))
}

func solve(input [9][9]byte) bool {
	// check rows
	for i := 0; i < 9; i++ {
		if !isValid(input[i]) {
			return false
		}
	}

	// make cols array
	var cols [9][9]byte
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cols[i][j] = input[j][i]
		}
	}

	// check cols
	for i := 0; i < 9; i++ {
		if !isValid(cols[i]) {
			return false
		}
	}

	// make squares
	var squares [9][9]byte
	squaresIdx := 0
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			elemIdx := 0
			square := [9]byte{}
			for y := i; y < i+3; y++ {
				for x := j; x < j+3; x++ {
					square[elemIdx] = input[y][x]
					elemIdx++
				}
			}
			squares[squaresIdx] = square
			squaresIdx++
		}
	}

	// check squares
	for i := 0; i < 9; i++ {
		if !isValid(squares[i]) {
			return false
		}
	}

	return true
}

func isValid(input [9]byte) bool {
	m := make(map[byte]int)
	for i := 0; i < 9; i++ {
		if input[i] == '.' {
			continue
		}
		if m[input[i]] > 0 {
			return false
		}
		m[input[i]]++
	}

	return true
}
