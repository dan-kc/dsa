package main

import (
	"fmt"
	"math"
)

func main() {
	// two crystal balls problem
	test := []bool{false, false, false, false, false, false, true}
	step := int(math.Sqrt(float64(len(test))))

	for i := step - 1; i < len(test); i += step {
		if test[i] {
			for j := i - step + 1; j <= len(test); j++ {
				if test[j] {
					fmt.Print(j)
					return
				}
			}
		}
	}
	// not found, meaning we go to last floor
	for j := len(test) - step; j < len(test); j++ {
		if test[j] {
			fmt.Print(j)
			return
		}
	}
}

// or

func TwoCrystlaBalls(a []bool) int {
	step := int(math.Sqrt(float64(len(a))))
	i := step - 1
	// incremint i until it's the last poss step
	for i < len(a) {
		if a[i] {
			break
		}
		i += step
	}
	// linear search remaining bit
	for j := i - step + 1; j < len(a); j++ {
		if a[j] {
			return j
		}
	}

	return -1
}
