import (
	"fmt"
)

func Permutations(n int) []int {
	if n < 4 {
		fmt.Println("NO SOLUTION")
		return
	}
	if n == 4 {
		fmt.Println("3 1 4 2")
		return
	}
	var odds []int
	var evens []int
	for i := n; i > 0; i-- {
		if i%2 == 0 {
			evens = append(evens, i)
		} else {
			odds = append(odds, i)
		}
	}
	all := append(evens, odds...)
	return all
}
