package misc

func ISqrt(x int) int {
	l := 0
	h := x + 1
	var m int
	for l != h-1 {
		m = l + (h-l)/2
		if m*m == x {
			return m
		}
		if m*m <= x {
			l = m
		} else {
			h = m
		}
	}
	return l
}
