package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	length := len(input)
	countMap := make(map[byte]int)
	for i := 0; i < length; i++ {
		countMap[input[i]]++
	}

	oddsCount := make(map[byte]int)
	evensCount := make(map[byte]int)
	for k, v := range countMap {
		if v%2 > 0 {
			oddsCount[k] = v
			continue
		}
		evensCount[k] = v
	}
	if len(oddsCount) > 1 {
		fmt.Print("NO SOLUTION")
		return
	}

	var res = make([]byte, length)
	for k, v := range oddsCount { // should only be one anyway
		for i := 0; i < v; i++ {
			res = append(res, k)
		}
	}
	for k, v := range evensCount {
		for i := 0; i < v/2; i++ {
			res = append([]byte{k}, res...) // prepend
			res = append(res, k)
		}
	}
	fmt.Print(string(res))
}
