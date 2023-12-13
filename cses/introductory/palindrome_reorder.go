package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	// count
	count := make(map[rune]int)
	for _, r := range input {
		count[r]++
	}

	// add odds first
	var res []rune
	var oddCount int
	for k, v := range count {
		if v%2 > 0 {
			oddCount++
			for i := 0; i < v; i++ {
				res = append(res, k)
			}
			count[k] = 0
		}
	}

	// no solution
	if oddCount > 1 {
		fmt.Println("NO SOLUTION")
		return
	}

	// add rest
	for k, v := range count {
		for i := 0; i < v; i += 2 {
			res = append(res, k)
			res = append([]rune{k}, res...)
		}
	}

	// print
	for _, r := range res {
		fmt.Printf("%c", r)
	}
}
