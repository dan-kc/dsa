package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	count := make(map[rune]int)
	for _, v := range str {
		count[v]++
	}
	var odds []rune
	for key, v := range count {
		if v%2 != 0 {
			odds = append(odds, key)
		}
	}
	if len(odds) > 1 {
		fmt.Print("NO SOLUTION")
		return
	}
	var ans string
	if len(odds) == 1 {
		ans = string(odds[0])
	}
	for key, v := range count {
		if v%2 == 0 {
			for i := 0; i < v/2; i++ {
				ans = string(key) + ans + string(key)
			}
		}
	}
	fmt.Println(ans)
}
