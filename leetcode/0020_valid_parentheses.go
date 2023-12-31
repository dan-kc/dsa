package main

import "fmt"

func main() {
	str := "(()[]{})"
	fmt.Print(solve(str))
}

func solve(str string) bool {
	stack := []byte{}
	brax := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for i := 0; i < len(str); i++ {
		_, ok := brax[str[i]]
		if ok {
			stack = append(stack, str[i])
			continue
		}

		topOfStack := stack[len(stack)-1]
		rbFromStack := brax[topOfStack]

		if rbFromStack != str[i] {
			return false
		}
		stack = stack[0 : len(stack)-1]
	}

	return len(stack) == 0
}
