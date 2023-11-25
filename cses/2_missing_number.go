package cses

func MissingNumber(n int, a []int) int {
	seen := make(map[int]bool)

	for i := 0; i < n-1; i++ {
		seen[a[i]] = true
	}

	for i := 1; i < n; i++ {
		if !seen[i] {
			return i
		}
	}

	return 0
}
