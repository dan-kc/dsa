package cses_test

import (
	"reflect"
	"testing"

	"github.com/dan-kc/dsa/cses"
)

func TestMissingNumber(t *testing.T) {
	n := 5
	a := []int{2, 3, 1, 5}
	expected := 4
	result := cses.MissingNumber(n, a)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For inputs %d and %d, expected %v, but got %v", n, a, expected, result)
	}
}
