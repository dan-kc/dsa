func IncreasingArray(n int, a []int) int {
	var moves int
	for i := 1; i < n; i++ {
		if a[i] < a[i-1] {
			moves = moves + a[i-1] - a[i]
			a[i] = a[i-1]
		}
	}
	return moves
}
