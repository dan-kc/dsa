package search

func BinarySearch(needle int, haystack []int) int {
	// l inclusive, h exclusive
	l := 0
	h := len(haystack)
	m := l + (h-l)/2

	for l < h {
		if needle == haystack[m] {
			return m
		} else if needle > haystack[m] {
			l = m + 1
			m = l + (h-l)/2
		} else {
			h = m - 1
			m = l + (h-l)/2
		}
	}

	return -1
}
