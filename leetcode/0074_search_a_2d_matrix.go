package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 4
	fmt.Print(solve(matrix, target))
}

func solve(matrix [][]int, target int) bool {
	ROWS := len(matrix)
	COLUMNS := len(matrix[0])
	l := 0
	r := ROWS

	for l < r {
		m := (l + r) / 2
		if target > matrix[m][len(matrix[m])-1] {
			l = m + 1
		} else if target < matrix[m][0] {
			r = m
		} else {
			break
		}
	}

	if l > r {
		return false
	}

	row := (l + r) / 2
	l = 0
	r = COLUMNS
	for l < r {
		m := (l + r) / 2
		if target > matrix[row][m] {
			l = m + 1
		} else if target < matrix[row][m] {
			r = m
		} else {
			return true
		}
	}

	return false
}
