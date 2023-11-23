package cses

func WeirdAlgorithm(n int) []int {
	var result []int

	for n != 1 {
		result = append(result, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}

	result = append(result, n)
	return result
}
