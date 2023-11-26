package misc

func TwoCrystlaBalls(a []bool) int {
	step := ISqrt(len(a))
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
