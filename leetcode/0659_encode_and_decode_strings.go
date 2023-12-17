package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := []string{"Daniel", "Cox"}
	encoded := encode(test)
	decoded := decode(encoded)
	fmt.Print(decoded)
}

func encode(input []string) string {
	var returnedString []rune
	for _, word := range input {
		runes := []rune(word)
		wordLen := []rune(strconv.Itoa(len(runes)))
		returnedString = append(returnedString, wordLen...)
		returnedString = append(returnedString, '#')
		returnedString = append(returnedString, []rune(word)...)
	}
	return string(returnedString)
}

func decode(input string) (res []string) {
	if len(input) == 0 {
		return
	}

	// convert to runes
	runes := []rune(input)

	// start cursor at 0
	var cursor int64 = 0
	for cursor < int64(len(runes)) {
		// stop cursor at '#' index
		var start = cursor
		for runes[cursor] != '#' {
			cursor++
		}

		// parse number before '#'
		lenString := runes[start:cursor]
		len, err := strconv.ParseInt(string(lenString), 10, 64)
		if err != nil {
			fmt.Print(err)
		}

		// add range into the res array
		word := runes[cursor+1 : cursor+1+len]
		res = append(res, string(word))

		// update cursor
		cursor = cursor + 1 + len
	}

	return
}
