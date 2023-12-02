package sort_test

import (
	"reflect"
	"testing"

	"github.com/dan-kc/dsa/sort"
)

func TestQuickSort(t *testing.T) {
	a := []int{2, 3, 1, 5}
	expected := []int{1, 2, 3, 5}
	result := sort.QuickSort(a)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For inputs %d, expected %v, but got %v", a, expected, result)
	}
}
