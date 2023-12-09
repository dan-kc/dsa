package leetcode

import (
	"reflect"
	"testing"
)

func ProductExceptSelf(nums []int) (res []int) {
	var pre, post []int

	for i, v := range nums {
		if i == 0 {
			pre = append(pre, i)
			continue
		}
		pre = append(pre, v*pre[i-1])
	}

	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			post = append([]int{nums[l-1]}, post...)
			continue
		}
		post = append([]int{nums[l] * pre[i+1]}, post...)
	}
	// [1,2,4,5]
	// pre [1,2,8,40]
	// post [40,40,20,5]

	pre = append([]int{1}, pre...)
	post = append(post, 1)

	for i := 0; i < l; i++ {
		res = append(res, pre[i-1]*post[i+1])
	}
	return
}

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 1}, []int{6, 3, 2, 6}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, v := range testCases {
		res := leetcode.ProductExceptSelf(v.input)
		if !reflect.DeepEqual(res, v.expected) {
			t.Errorf("Input: %v, Expected: %v, Result: %v", v.input, v.expected, res)
		}
	}
}
