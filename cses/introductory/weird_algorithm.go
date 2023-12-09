import (
	"reflect"
	"testing"
)

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

	return append(result, n)
}

func TestWeirdAlgorithm(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{3, []int{3, 10, 5, 16, 8, 4, 2, 1}},
	}

	for _, tc := range testCases {
		result := cses.WeirdAlgorithm(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For input %d, expected %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
