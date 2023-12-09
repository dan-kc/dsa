func Repetitions(sequence string) int {
	maxRepetition := 1
	currentRepetition := 1
	for i := 1; i < len(sequence); i++ {
		if sequence[i] == sequence[i-1] {
			currentRepetition++
		} else {
			if currentRepetition > maxRepetition {
				maxRepetition = currentRepetition
			}
			currentRepetition = 1
		}
	}
	if currentRepetition > maxRepetition {
		maxRepetition = currentRepetition
	}

	return maxRepetition
}
