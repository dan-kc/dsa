func TwoKnights(n int) (res []int) {
	res = append(res, 0)
	for k := 2; k < n+1; k++ {
		totalMoves := k * k * (k*k - 1) / 2
		attackMoves := (k - 2) * (k - 1)
		res = res.append(res, totalMoves-(4*attackMoves))
	}
}
