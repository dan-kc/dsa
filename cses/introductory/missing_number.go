import (
	"reflect"
	"testing"
)

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

func TestMissingNumber(t *testing.T) {
	n := 5
	a := []int{2, 3, 1, 5}
	expected := 4
	result := cses.MissingNumber(n, a)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For inputs %d and %d, expected %v, but got %v", n, a, expected, result)
	}
}
