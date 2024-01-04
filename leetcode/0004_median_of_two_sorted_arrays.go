package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2}
	arr2 := []int{3, 4}
	fmt.Println(solve(arr, arr2))
}

func solve(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(nums1) + len(nums2)
	half := (total + 1) / 2

	var Aleft, Aright, Bleft, Bright int

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)-1
	for {
		i := (l + r) >> 1 // A
		j := half - i - 2 // B

		if i >= 0 {
			Aleft = A[i]
		} else {
			Aleft = math.MinInt
		}

		if (i + 1) < len(A) {
			Aright = A[i+1]
		} else {
			Aright = math.MaxInt
		}

		if j >= 0 {
			Bleft = B[j]
		} else {
			Bleft = math.MinInt
		}

		if (j + 1) < len(B) {
			Bright = B[j+1]
		} else {
			Bright = math.MaxInt
		}

		// partition is correct
		if Aleft <= Bright && Bleft <= Aright {
			// odd
			if total%2 == 1 {
				return float64(max(Aleft, Bleft))
			}
			// even
			return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
