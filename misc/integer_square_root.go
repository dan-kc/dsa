package misc

// Time: O(log(N))
// Space: O(1)
func ISqrt(x int) int {
	l := 0
	h := x + 1
	var m int
	for l != h-1 {
		m = l + (h-l)/2
		if m*m <= x {
			l = m
		} else {
			h = m
		}
	}
	return l
}
